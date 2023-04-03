package main

import "fmt"

func main() {
	var slice = []string{"João", "Joanilton", "Ramires", "Samanta", "João", "Lucas", "Alice", "Jumento"}
	var nome string
	fmt.Print("digite um nome: ")
	fmt.Scan(&nome)
	novoSlice := []string{}
	for _, element := range slice {
		if element != nome {
			novoSlice = append(novoSlice, element)
		}
	}
	fmt.Println(novoSlice)
}
