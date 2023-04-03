package main

import "fmt"

func main() {
	var matriz = [10]int{2, 6, 7, 34, 7, 2, 34, 65, 7, 2}
	index := -1
	x := 0
	for _, element := range matriz {
		index++
		if index%2 == 0 {
			x += element
		}
	}
	fmt.Println(x)
}
