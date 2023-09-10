//Crie um Array de inteiros com 10 elementos. Em seguida, solicite ao usuário
//que informe um valor e verifique se esse valor está presente no Array.
//Imprima o resultado da busca.

package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var n1 int
	fmt.Print("Informe um valor: ")
	fmt.Scan(&n1)

	for _, elementoarray := range array {
		if elementoarray == n1 {
			fmt.Println("O número está dentro da array.")
			return
		}

	}
	fmt.Print("O número não está dentro do array.")
}
