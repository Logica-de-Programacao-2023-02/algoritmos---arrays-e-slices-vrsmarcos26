// Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5. Remova o terceiro elemento e imprima o Slice resultante.

package main

import "fmt"

func main() {
	nin := []int{1, 2, 3, 4, 5}
	nin = append(nin[:2], nin[3:]...)
	fmt.Println(nin)
}
