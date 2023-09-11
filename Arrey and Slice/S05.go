// Crie um Array bidimensional de inteiros com 3 linhas e 2 colunas.
// Solicite ao usuário que informe os valores de cada elemento da matriz.
// Em seguida, imprima a matriz resultante.
package main

import (
	"fmt"
)

func main() {
	var matriz [3][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Informe o valor da da linha[%d] coluna [%d]: ", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}

	for _, ints := range matriz {
		fmt.Print(ints)
	}

}
