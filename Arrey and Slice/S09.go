// Crie um Array de floats com 6 elementos. Solicite ao usuário que informe
// um número que será adicionado em todas as posições do Array. Imprima o Array resultante.
package main

import "fmt"

func main() {
	array := [6]float64{1, 2, 3, 4, 5, 6}
	fmt.Print("Informe um número: ")
	var n1 float64
	fmt.Scan(&n1)
	for i := 0; i < len(array); i++ {
		array[i] += n1
	}

	fmt.Print(array)

}
