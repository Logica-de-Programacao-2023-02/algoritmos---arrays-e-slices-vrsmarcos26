//Crie um Array de inteiros com 7 elementos. Solicite ao usuário que informe um número que será
//adicionado ao primeiro e ao último elemento do Array. Imprima o Array resultante.

package main

import "fmt"

func main() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}

	var n, n2 int
	fmt.Print("Informe um número que será adicionado em primeiro: ")
	fmt.Scan(&n)
	fmt.Print("Informe um número que será adicionado em último: ")
	fmt.Scan(&n2)

	novoArray := [9]int{n}
	copy(novoArray[1:8], array[:])
	novoArray[8] = n2

	fmt.Println("Array resultante:", novoArray)
}
