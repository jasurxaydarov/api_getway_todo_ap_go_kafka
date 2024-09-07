package api

import (
	"github.com/jasurxaydarov/api_getway_todo_ap_go_kafka/api/handlers"
	"github.com/jasurxaydarov/api_getway_todo_ap_go_kafka/event"
	"github.com/labstack/echo/v4"
)

type Options struct {
	Event event.EventI
}


func Api(o Options)*echo.Echo{
	h:=handlers.NewHandler(o.Event)

	router:=echo.New()

	api:=router.Group("/api")

	{
		api.POST("/create-todo",h.CreateTodo)
		api.POST("/update-todo/:id",h.UpdateTodo)

	}	

	return router
}