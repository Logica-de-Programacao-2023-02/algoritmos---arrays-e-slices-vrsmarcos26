// Crie um Slice de inteiros com tamanho 10 e imprima o valor mínimo e o valor
// máximo armazenados no Slice.
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	vmax := slice[0]
	for _, elementoarray := range slice {
		if elementoarray > vmax {
			vmax = elementoarray
		}
	}

	vmin := slice[0]
	for _, elementoarray := range slice {
		if elementoarray < vmin {
			vmin = elementoarray
		}
	}

	fmt.Printf("O valor mínimo é %d.\n", vmin)
	fmt.Printf("O valor mínimo é %d.\n", vmax)

}
