// Crie um Slice de inteiros e solicite ao usuário que informe o tamanho
//do Slice e os valores dos elementos. Em seguida, imprima o Slice e a
//soma dos valores armazenados.

package main

import "fmt"

func main() {
	var t int
	fmt.Print("Tamanho da Slice que você quer: ")
	fmt.Scan(&t)

	Slice := make([]int, t)

	for i := 0; i < t; i++ {
		fmt.Printf("Valor dos elesmentos %d: ", i+1)
		fmt.Scan(&Slice[i])
	}
	fmt.Println("Essa é sua Slice de números int: ", Slice)

	soma := 0
	for _, valor := range Slice {
		soma += valor
	}

	fmt.Print("O valor da soma dos seus elementos na slice é: ", soma)

}
