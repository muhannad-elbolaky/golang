package models

import (
	"fmt"
	"time"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (event *Event) Save() {
	// TODO: save to database
	fmt.Println(*event)
	events = append(events, *event)
}

func GetAllEvents() []Event {
	return events
}
