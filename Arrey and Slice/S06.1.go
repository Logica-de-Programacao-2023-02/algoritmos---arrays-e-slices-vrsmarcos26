package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var n1 int
	fmt.Print("Digite um valor para verificar se está presente no array: ")
	fmt.Scan(&n1)

	var achei bool

	for _, elementoarray := range array {
		if elementoarray == n1 {
			achei = true
		}
	}

	if achei {
		fmt.Println("O valor está presente no array.")
	} else {
		fmt.Println("O valor não está presente no array.")
	}
}
