package handlers

import "github.com/jasurxaydarov/api_getway_todo_ap_go_kafka/event"

type handler struct {
	event event.EventI
}

func NewHandler(event event.EventI) *handler {

	return &handler{event: event}
}


