package main

import "fmt"

func main() {
	var n int

	fmt.Print("Informe o tamanho das slices: ")
	fmt.Scan(&n)

	slice1 := make([]int, n)
	slice2 := make([]int, n)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Print("Digite o valor da slice 1: ")
		fmt.Scan(&slice1[i])
	}

	for i := 0; i < n; i++ {
		fmt.Print("Digite o valor da slice 2: ")
		fmt.Scan(&slice2[i])
	}

	for i := 0; i < n; i++ {
		result[i] = slice1[i] + slice2[i]
	}
	fmt.Print(result)

}
