// Crie um Array bidimensional de inteiros com 3 linhas e 2 colunas.
// Solicite ao usuário que informe os valores de cada elemento da matriz.
// Em seguida, imprima a matriz resultante.
package main

import "fmt"

func main() {
	var matriz [3][2]int

	for i := 0; i < 3; i++ {
		for n := 0; n < 2; n++ {
			fmt.Printf("Informe o valor da matriz[%d][%d]: ", i, n)
			fmt.Scan(&matriz[i][n])
		}
	}
	fmt.Println("Matriz resultante:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Println()
	}

}
