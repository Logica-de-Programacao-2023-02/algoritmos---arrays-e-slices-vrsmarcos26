//Crie um Array de inteiros com 10 elementos.
//Calcule e imprima a soma dos elementos nas posições pares do Array.

package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	soma := 0
	for _, n := range array {
		if n%2 == 0 {
			soma += n
		}
	}
	fmt.Print(soma)
}
