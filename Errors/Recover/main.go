package main

import (
	"fmt"
)

func dividir(num1 int, num2 int) int {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recover: ", r)
		}
	}()

	return num1 / num2
}

func dividir2(num1 int, num2 int) int {
	return num1 / num2
}

func main() {
	dividir(10, 0)
	fmt.Println("Sem o recover isso não aparece")
	dividir2(20, 0)
	fmt.Println("Viu só")
}
