package main

import (
	"fmt"
)

func main() {
	nums := []int{1}
	// fmt.Println(secondToLast(nums))

	fmt.Println(safeSecondToLast(nums))
}

func safeSecondToLast(nums []int) (i int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	return secondToLast(nums), nil

}

func secondToLast(nums []int) int {
	return nums[len(nums)-2]
}
