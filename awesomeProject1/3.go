package main

import (
	"fmt"
)

func main() {
	var matriz = [4]float64{6.8, 7.6, 5.5, 6.7}
	var produto float64
	produto = 1
	for _, element := range matriz {
		produto *= element

	}

	fmt.Println(produto)
}
