package main

import (
	"fmt"
)

func main() {
	var array = [5]int{3, 5, 15, 56, 71}
	var slice = []int{}
	var element int
	for _, element = range array {
		if element%3 == 0 {
			slice = append(slice, element)
		}

	}
	fmt.Println(slice)

}
