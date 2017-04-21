package event

import (
	"errors"
	"log"
)

// Templates containes sample events as a preview to the user
var Templates map[string]*Event

// Events contains the registered events in the application
var Events map[string]*Event

func init() {
	Templates = make(map[string]*Event)
	Events = make(map[string]*Event)
}

// Register registers events as templates for the user
func Register(trigger Trigger) {
	newEvent := New(trigger)
	Templates[newEvent.Type()] = newEvent
}

// AddEvent adds an event to the application
func AddEvent(event *Event) error {
	_, found := Events[event.ID()]
	if found {
		return errors.New("Event with that id already exists")
	}
	Events[event.ID()] = event

	err := event.Listen()
	if err != nil {
		log.Println(err)
	}
	return nil
}

// RemoveEvent removes an evenet from the application
func RemoveEvent(event *Event) {
	delete(Events, event.ID())
}
