package main

import (
	"fmt"
)

func main() {
	multiply := func(number []int, factor int) []int {
		res := make([]int, len(number))
		for i, e := range number {
			res[i] = e * factor
		}
		return res
	}
	add := func(number []int, factor int) []int {
		res := make([]int, len(number))
		for i, e := range number {
			res[i] = e + factor
		}
		return res
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range add(multiply(ints, 10), 1) {
		fmt.Printf("%d ", v)
	}

}
