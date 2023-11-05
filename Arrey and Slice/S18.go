//Escreva um programa que leia um número inteiro positivo n e gere um
//array com os n primeiros números primos.

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("O número deve ser um inteiro positivo.")
		return
	}

	primeNumbers := make([]int, 0)
	num := 2

	for len(primeNumbers) < n {
		if isPrime(num) {
			primeNumbers = append(primeNumbers, num)
		}
		num++
	}

	fmt.Println("Os", n, "primeiros números primos são:", primeNumbers)
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
