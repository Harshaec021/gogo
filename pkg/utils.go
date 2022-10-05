package utils

import (
	"fmt"
)

func mean(num []int) float64 {
	var mean_int float64
	fmt.Println("About to find sum")
	s := sum(num)
	fmt.Println(s)

	mean_int = float64(s) / float64(len(num))

	return mean_int
}

func sum(nums []int) int {

	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}
