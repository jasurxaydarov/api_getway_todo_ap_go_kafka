package event

import (
	"encoding/json"
	"log"

	"github.com/jasurxaydarov/api_getway_todo_ap_go_kafka/dto"
	"github.com/jasurxaydarov/api_getway_todo_ap_go_kafka/models"
)


type EventI interface{
	CreateTodo(data *dto.CreateTodoDto)error
	UpdateTodo(data *models.TodoModel)error
}

type event struct{

}

func NewEvent()EventI{

	return &event{}
}


func (e *event)CreateTodo(data *dto.CreateTodoDto)error{

	topic:="create-todo"

	msg,err:=json.Marshal(data)

	if err!=nil{

		log.Println("err on json.Marshal ", err)
		return err
	}


	err=SendMessage(topic,msg)
	if err!=nil{

		log.Println("err on err=SendMessage(topic,msg)", err)
		return err
	}

	return nil
}


func (e *event)UpdateTodo(data *models.TodoModel)error{

	topic:="update-todo"

	msg,err:=json.Marshal(data)

	if err!=nil{

		log.Println("err on json.Marshal ", err)
		return err
	}


	err=SendMessage(topic,msg)
	if err!=nil{

		log.Println("err on err=SendMessage(topic,msg)", err)
		return err
	}

	return nil
}