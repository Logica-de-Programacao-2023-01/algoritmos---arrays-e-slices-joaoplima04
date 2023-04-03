package main

import "fmt"

func main() {
	var slice = []int{2, 4, 6, 7, 8, 4, 34, 62}
	var index, indexi int
	fmt.Println("Digite o índice do primeiro número que deseje trocar de posição:  ")
	fmt.Scan(&indexi)
	fmt.Println("Digite o segundo: ")
	fmt.Scan(&index)
	slice[index], slice[indexi] = slice[indexi], slice[index]
	fmt.Println(slice)
}
