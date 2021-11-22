package main

import (
	"fmt"
)

func filter(pred func(int) bool, values []int) []int {
	var odds []int
	for value := range values {
		if pred(value) {
			odds = append(odds, value)

		} else {
			fmt.Printf("%v is not odd \n", value)
		}
	}
	return odds
}

func is_odd(n int) bool {
	return n%2 == 1
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Printf("odd values are %v", filter(is_odd, nums))

}
