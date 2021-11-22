package main

import (
	"fmt"
)

func main() {
	fmt.Println("mean:", mean([]int{1,2,3}))
	fmt.Println("mean:",mean([]int{1,2,3,4}))


}


func mean(nums []int ) float64 {
		s := sum(nums)
		mean := float64(s) / float64(len(nums))
		return mean
}

func sum(nums []int) int {
		total := 0
		for _, n := range nums {
			total += n
		}
		fmt.Println("total:",total)

		return total
}