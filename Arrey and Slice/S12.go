// Crie um Array de inteiros com 5 elementos. Em seguida, crie um novo Slice que
// contenha apenas os elementos do Array que são múltiplos de 3. Imprima o novo Slice.
package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := []int{}

	for _, n := range array {
		if n%3 == 0 {
			slice = append(slice, n)
		}
	}
	fmt.Print(slice)
}
