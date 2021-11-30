package main 

import (
		"fmt"
		"log"
		"time"
)

type Event struct {
	ID string
	Time time.Time
}



type DoorEvent struct {
	Event 
	Action string
}

type TemparatureEvent struct {
	Event 
	Value float64
}

func NewDoorEvent(id string, time time.Time, action string) (*DoorEvent, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id")
	}
	evt := DoorEvent{
		Event: Event{id, time},
		Action: action,
	}
	return &evt, nil
}

func main() {
	evt, err := NewDoorEvent("front door", time.Now(),"open")
	if err !=nil {
			log.Fatal(err)
	}
	fmt.Printf("%+v\n",evt)
}