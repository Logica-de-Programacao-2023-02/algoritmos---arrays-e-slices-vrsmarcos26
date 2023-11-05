//Crie um Slice de inteiros com tamanho 8 e solicite ao usuário que informe dois índices de
//elementos que deverão ser trocados de posição. Imprima o Slice resultante.

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Print("Obeserve essa slice que você tem. \n ", slice)

	var n, n2 int

	fmt.Print("Informe um indice: ")
	fmt.Scan(&n)
	fmt.Print("Informe outro indice: ")
	fmt.Scan(&n2)

	if n < 0 || n >= len(slice) || n2 < 0 || n2 >= len(slice) {
		fmt.Println("Os índices informados estão fora dos limites da slice.")
		return
	}

	slice[n], slice[n2] = slice[n2], slice[n]
	fmt.Println(slice)
}
