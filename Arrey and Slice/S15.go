//Crie um Array de floats com 10 elementos. Crie um novo Slice que contenha apenas os elementos do
//Array que são maiores que 5. Imprima o novo Slice.

package main

import "fmt"

func main() {
	array := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice := []float64{}

	for _, n := range array {
		if n > 5 {
			slice = append(slice, n)
		}
	}

	fmt.Print(slice)
}
