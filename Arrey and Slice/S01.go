package main

import "fmt"

func main() {
	var numerosin = [3]int{1, 2, 3}
	var soma int
	for _, num := range numerosin {
		soma += num
	}
	fmt.Print(soma)
}
