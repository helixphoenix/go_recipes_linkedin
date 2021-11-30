package main 

import (
	 	"fmt"
		 "log"
)

type ClickEvent struct {

}


type HoverEvent struct {

}

var eventCounts = make(map[string]int)


func recordEvent(evt interface{}) {
	switch evt.(type){
	case *ClickEvent:
			eventCounts["click"]++
	case *HoverEvent:
			eventCounts["hover"]++

	default:
			log.Printf("WARNING: unknown events are happening: %#v of type %T\n", evt, evt)

	}

}


func main() {
	recordEvent(&ClickEvent{})
	recordEvent(&HoverEvent{})
	recordEvent(&ClickEvent{})
	recordEvent("Various hacking events")

	fmt.Println("event counts:", eventCounts)


}