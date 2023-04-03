package main

import "fmt"

func main() {
	var array = [10]float64{6.8, 4.3, 5.4, 2, 6.8, 4.8, 5, 9, 1.3, 4}
	var slice = []float64{}
	for _, valor := range array {
		if valor > 5 {
			slice = append(slice, valor)
		}
	}
	fmt.Println(slice)

}
