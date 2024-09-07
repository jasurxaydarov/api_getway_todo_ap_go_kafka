package event

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

const KafkaBroker = "localhost:9092"

var writer *kafka.Writer

func initKafkaWriter(topic string) *kafka.Writer {

	kafkaWriter := kafka.NewWriter(
		kafka.WriterConfig{
			Brokers:  []string{KafkaBroker},
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	)

	return kafkaWriter
}

func SendMessage(topic string, msg []byte) error {

	if writer == nil {
		writer = initKafkaWriter(topic)
	}

	if err := writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: msg,
		},
	); err != nil {

		log.Println("err on  writer.WriteMessages() Error: ", err)

		return err
	}

	return nil
}

func CloseKafkaWriter() error {

	if writer != nil {
		if err := writer.Close(); err != nil {

			log.Println("err on writer.Close()", err)
			return err
		}

		writer = nil
	}

	return nil

}
