package event

import "github.com/AUT-Daralfonoon/back-end/modules/holding/event"

type EventController struct{
	EventManger *event.EventManager
}

func NewEventController(em *event.EventManager) *EventController{
	return &EventController{em}
}
