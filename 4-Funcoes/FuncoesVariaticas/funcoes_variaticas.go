package main

import "fmt"

//Recebe infinitos valores int
func soma(numeros ...int) int {
	// O valor vem como um slice
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total

}

func main() {
	soma1 := soma(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(soma1)

	soma2 := soma(484, 0552454, 07477, 7757575587, 7858778)
	fmt.Println(soma2)

	soma3 := soma()
	fmt.Println(soma3)

}
