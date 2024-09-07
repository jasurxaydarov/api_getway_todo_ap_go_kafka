package main

import (
	"github.com/jasurxaydarov/api_getway_todo_ap_go_kafka/api"
	"github.com/jasurxaydarov/api_getway_todo_ap_go_kafka/event"
)

func main(){
	event:=event.NewEvent()

	router:=api.Api(
		api.Options{
			Event: event,
		},
	)

	router.Start(":8080")
	
}


