package main

import "fmt"

func main() {
	var numeros = [3]int{1, 2, 3}
	var soma int
	for _, numeros := range numeros {
		soma += numeros
	}
	fmt.Println(soma)
}
