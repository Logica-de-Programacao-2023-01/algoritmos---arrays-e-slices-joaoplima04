package main

import "fmt"

func main() {
	var array = [7]int{1, 3, 5, 6, 2, 4, 5}
	var elemento, elementu int
	fmt.Print("Digite um número que será adicionado ao primeiro elemento: ")
	fmt.Scan(&elemento)
	array[0] = elemento
	fmt.Print("Digite um número que será adicionado ao último elemento: ")
	fmt.Scan(&elementu)
	array[6] = elementu
	fmt.Println(array)

}
