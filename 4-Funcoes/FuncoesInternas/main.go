package main

import "fmt"

func main() {
	// A partir do go 1.21
	maximo := min(1, 2, 3, 588, 555) // retorna o valor minimo
	fmt.Println(maximo)

	minimo := max(342523, 324, 232, 32) // retorna o valor maximo
	fmt.Println(minimo)

}
