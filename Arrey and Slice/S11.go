// Crie um Array bidimensional de inteiros com 2 linhas e 3 colunas.
// Em seguida, solicite ao usuário que informe um índice de linha e outro de coluna
// e imprima o valor armazenado nessa posição da matriz.
package main

import "fmt"

func main() {
	matriz := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Matriz:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Println()
	}
	var linha, coluna int
	fmt.Println("Informe um índice de linha: ")
	fmt.Scan(&linha)
	fmt.Println("Informe um indice de coluna: ")
	fmt.Scan(&coluna)

	if linha >= 0 && linha < 2 && coluna >= 0 && coluna < 3 {
		valor := matriz[linha][coluna]
		fmt.Printf("O valor na posição (%d, %d) é: %d\n", linha, coluna, valor)
	} else {
		fmt.Println("Índices fora do intervalo da matriz.")
	}
}
