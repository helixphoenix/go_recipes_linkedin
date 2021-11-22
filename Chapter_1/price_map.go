package main

import (
	"fmt"
)


func main(){
	okCheck()
}

func okCheck() {
	prices := map[string]int{
		"Banana": 0,
	}

	banana, ok := prices["Banana"]
	if ok {
		fmt.Printf("The price of the Banana is  $%d\n", banana)
	} else {
		fmt.Printf("no banana no cry")
	}

	apple, ok := prices["Apple"]
	if ok {
		fmt.Printf("The price of the Apple is $%d\n", apple)
	} else {
		fmt.Println("no apple no cry")
	}

	dudu, ok := prices["Dudu"]
	if ok {
		fmt.Printf("The price of the Apple is $%d\n", dudu)
	} else {
		fmt.Println("no dudu no cry")
	}

}
