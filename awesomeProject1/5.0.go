package main

import "fmt"

func main() {
	const linhas, colunas = 3, 2
	matriz := [linhas][colunas]int{}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			fmt.Printf("Digite o valor para o elemento [%d][%d]: ", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}

	fmt.Println("Matriz resultante:")
	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Println()
	}
}
