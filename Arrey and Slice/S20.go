package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	cres := true

	for i := 0; i < len(arr)-1; i++ {
		atl := arr[i]
		prox := arr[i+1]
		if atl >= prox {
			cres = false
			break
		}
	}

	if cres {
		fmt.Print("array esá ordenado.")
	} else {
		fmt.Print("não está ordenado")
	}
}
