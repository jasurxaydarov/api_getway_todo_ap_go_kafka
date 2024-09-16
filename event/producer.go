package event

import (
	"encoding/json"
	"log"

	"github.com/jasurxaydarov/api_getway_todo_ap_go_kafka/models"
)

type EventI interface {
	CreateTodo(data *models.NewTodo) error
	UpdateTodo(data *models.UpdateTodo) error
	DeleteTodo(data *models.DeleteByID) error
}

type event struct {
}

func NewEvent() EventI {

	return &event{}
}

func (e *event) CreateTodo(data *models.NewTodo) error {

	topic := "create-todo"

	msg, err := json.Marshal(data)

	if err != nil {

		log.Println("err on json.Marshal ", err)
		return err
	}

	err = SendMessage(topic, msg)
	if err != nil {

		log.Println("err on err=SendMessage(topic,msg)", err)
		return err
	}

	return nil
}

func (e *event) UpdateTodo(data *models.UpdateTodo) error {

	topic := "update-todo"

	msg, err := json.Marshal(data)

	if err != nil {

		log.Println("err on json.Marshal ", err)
		return err
	}

	err = SendMessage(topic, msg)
	if err != nil {

		log.Println("err on err=SendMessage(topic,msg)", err)
		return err
	}

	return nil
}

func (e *event) DeleteTodo(data *models.DeleteByID) error {

	topic := "delete-todo"

	msg, err := json.Marshal(data)

	if err != nil {

		log.Println("err on json.Marshal ", err)
		return err
	}

	err = SendMessage(topic, msg)
	if err != nil {

		log.Println("err on err=SendMessage(topic,msg)", err)
		return err
	}

	return nil
}
