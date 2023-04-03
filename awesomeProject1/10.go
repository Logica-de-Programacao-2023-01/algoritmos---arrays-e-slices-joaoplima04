package main

import "fmt"

func main() {
	var slice = []int{32, 76, 54, 60, 21, 87, 33, 56, 23, 54}
	max := 0
	min := 10000
	for _, element := range slice {
		if element > max {
			max = element
		}
		if element < min {
			min = element
		}
	}
	fmt.Println("O valor mínimo é:", min)
	fmt.Println("O valor máximo é:", max)
}
