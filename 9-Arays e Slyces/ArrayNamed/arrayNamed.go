package main

import (
	"fmt"
)

func main() {
	type numeros [3]int
	type numbers [3]int

	meuArray := [3]int{1, 2, 3}
	meusNumeros := numeros{1, 2, 3}
	myNumbers := numbers{1, 2, 3}

	fmt.Printf("meuArray == meusNumeros %v\n", meuArray == meusNumeros)
	fmt.Printf("meuArray == myNumbers %v\n", meuArray == myNumbers)
	// fmt.Printf("meusNumeros == myNumbers %v\n", meusNumeros == myNumbers)
	fmt.Println("Não posso comparar meusNumeros == myNumbers pois os tipos são diferentes")

	fmt.Printf("numeros(myNumbers) == meusNumeros %v\n", numeros(myNumbers) == meusNumeros)

}
