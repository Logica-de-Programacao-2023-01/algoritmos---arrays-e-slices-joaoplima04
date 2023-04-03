package main

import "fmt"

func main() {
	var array = [10]int{1, 3, 8, 9, 16, 5, 6, 7, 4, 20}
	var valor int
	fmt.Print("Digite um valor: ")
	fmt.Scan(&valor)
	for _, elemento := range array {
		if elemento == valor {
			fmt.Println("este valor está presente na matriz")
		} else {
			fmt.Println("Valor não está presente na matriz")
			break
		}

	}
}
