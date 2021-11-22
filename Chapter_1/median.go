package main

import (
	"fmt"
	"sort"
)


func main (){
		fmt.Println(median([]float64{3,1,2}))
		fmt.Println(median([]float64{3,1,4,2}))

}

func median(nums []float64) float64 {
	    // memory allocation for copying nums
		vals := make([]float64, len(nums))
		//copying nums to vals
		copy(vals, nums)
		fmt.Println("values sliced:", vals)
		sort.Float64s(vals)
		i := len(vals)/2
		if len(vals)%2 == 1 {
			return vals[i]
		}
		return (vals[i-1] + vals[i]) / 2
}