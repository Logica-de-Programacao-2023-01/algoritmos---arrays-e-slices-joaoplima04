package main

import "fmt"

func main() {
	var tamanho, x, soma int
	fmt.Println("Qual o tamanho do slice? ")
	fmt.Scan(&tamanho)
	slice := make([]int, tamanho)
	x = 0
	for x < tamanho {
		x++
		fmt.Println("Digite o valor para o elemento:", x)
		fmt.Scan(&slice[x])
	}
	for _, element := range slice {
		soma += element
	}
	fmt.Println(slice)
	fmt.Println("A soma dos elementos é: ", soma)
}
