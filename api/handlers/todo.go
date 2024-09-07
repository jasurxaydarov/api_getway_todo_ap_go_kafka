package handlers

import (
	"log"

	"github.com/google/uuid"
	"github.com/jasurxaydarov/api_getway_todo_ap_go_kafka/dto"
	"github.com/jasurxaydarov/api_getway_todo_ap_go_kafka/models"
	"github.com/labstack/echo/v4"
)

func (h *handler) CreateTodo(c echo.Context) error {

	reqBody := dto.CreateTodoDto{}

	if err := c.Bind(&reqBody); err != nil {

		log.Println("err on CreateTodo  c.Bind", err)
		return err
	}

	if err := h.event.CreateTodo(&reqBody); err != nil {
		log.Println("err on CreateTodo  h.event.CreateTodo", err)
		return err

	}

	c.JSON(201,reqBody)
	return nil

}

func (h *handler) UpdateTodo(c echo.Context) error {

	reqBody := models.TodoModel{}

	
	
	id,_:=uuid.Parse(c.Param("id"))

	reqBody.Id=id
	if err := c.Bind(&reqBody); err != nil {

		log.Println("err on updateTodo  c.Bind", err)
		return err
	}

	if err := h.event.UpdateTodo(&reqBody); err != nil {
		log.Println("err on update  h.event.CreateTodo", err)
		return err

	}

	c.JSON(201,reqBody)
	return nil

}

