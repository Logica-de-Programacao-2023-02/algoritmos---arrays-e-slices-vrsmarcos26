package main

import "fmt"

func main() {
	var numerosin = [4]float64{1, 2, 3, 4}
	var mult float64 = 1
	for _, num := range numerosin {
		mult *= num
	}
	fmt.Print(mult)
}
