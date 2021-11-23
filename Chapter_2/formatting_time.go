package main 

import (
	"fmt"
	"time"
)

func main() {
	lennon := time.Date(1940, time.October, 9, 18, 30, 0, 0, time.UTC)
	fmt.Println(lennon, lennon.Weekday())

	fmt.Println(lennon.Format("02-01-2006"))
	fmt.Println(lennon.Format("Monday, Jan 02"))

	fmt.Println(lennon.Format(time.RFC3339Nano))
   
    d := 3500 * time.Millisecond
	fmt.Println(d) 
}