// // convert time stamp from one zone to another

package main

import (
	"fmt"
	"time"
)

func tsConvert(ts, convert_from, convert_to string) (string, error) {
	from, err := time.LoadLocation((convert_from))
	if err != nil {
		fmt.Printf("Wrong location format %v", err)
	}

	const format = "2006-01-02T15:04"
	from_time,err := time.ParseInLocation(format, ts, from )
	if err != nil {
		fmt.Printf("Due to %v it did not work ", err)
	}
	fmt.Printf("the time for %v is %s \n", convert_from, from_time)



	to, err := time.LoadLocation((convert_to))
	if err != nil {
		fmt.Printf("Wrong location format %v", err)
	}

	to_time := from_time.In(to)

	return to_time.Format(format), err
}





func main() {
	ts := "2021-03-08T19:12"
	out, err := tsConvert(ts, "America/Los_Angeles", "Asia/Jerusalem")
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	fmt.Println( out) // 2021-03-09T05:12
}



