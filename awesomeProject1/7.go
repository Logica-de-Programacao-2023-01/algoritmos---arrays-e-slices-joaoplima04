package main

import (
	"fmt"
)

func main() {
	var slice = []int{2, 6, 9, 3, 1}
	var numero, elemento int
	fmt.Println("Slice: ", slice)
	fmt.Print("Digite um número inteiro para adicionar ao slice: ")
	fmt.Scan(&numero)
	for _, elemento = range slice {
		if numero == elemento {
			fmt.Println("Número já está na lista")
			break
		}
	}
	if numero != elemento {
		slice = append(slice, numero)
		fmt.Println(slice)
		fmt.Println("Número adicionado com sucesso!")
	}

}
