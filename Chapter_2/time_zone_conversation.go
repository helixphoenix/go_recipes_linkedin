package main 

import (
	"fmt"
	"time"
)





func main() {
	chi, err := time.LoadLocation(("America/Chicago"))
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	chiTime := time.Date(2021, time.February, 29, 19, 30, 0 , 0 , chi)
	fmt.Println("Chicago:",chiTime)
    
	ist, err := time.LoadLocation(("Asia/Istanbul"))
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	istTime := chiTime.In(ist)
	fmt.Println("IZM:", istTime)

  
}