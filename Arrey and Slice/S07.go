//Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que
//informe um número inteiro. Adicione esse número ao Slice apenas se ele não estiver presente.
//Imprima o Slice resultante.

package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}

	var n1 int
	fmt.Print("Informe um número: ")
	fmt.Scan(&n1)

	for _, elementoarray := range slice {
		if elementoarray == n1 {
			fmt.Println(slice)
			return
		}

	}
	for _, elementoarray := range slice {
		if elementoarray != n1 {
			slice = append(slice, n1)
			fmt.Println(slice)
			return
		}

	}
}
