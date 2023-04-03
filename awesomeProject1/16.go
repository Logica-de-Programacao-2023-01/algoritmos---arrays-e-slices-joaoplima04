package main

import "fmt"

func main() {
	var array = [10]int{26, 46, 345, 6, 867, 34, 234, 76, 3, 7}
	var slice = []int{}
	for _, element := range array {
		if element%2 == 0 {
			slice = append(slice, element)
		}
	}
	fmt.Println(slice)

}
