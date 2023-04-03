package main

import "fmt"

func main() {
	var array = [6]float64{6.8, 5.6, 3.4, 6.6, 6.7, 4.6}
	var valor float64
	fmt.Println("Digite um valor decimal: ")
	fmt.Scan(&valor)
	for x := 0; x < len(array); x++ {
		array[x] += valor
	}
	fmt.Println(array)
}
