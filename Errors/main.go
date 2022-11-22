package main

import (
	"fmt"
)

type MathErr string

func (me MathErr) Error() string {
	return string(me)
}

// Sentinel Errors
const (
	ErrDivisaoPorZero = MathErr("Não é possível dividir por 0")
	ErrNumeroNegativo = MathErr("Não é possível dividir um numero negativo")
)

func dividir(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, ErrDivisaoPorZero
	}
	if num1 < 0 {
		return 0, ErrNumeroNegativo
	}
	return num1 / num2, nil
}

func main() {
	_, err := dividir(10, 0)
	if err != nil {
		if err == ErrDivisaoPorZero {
			fmt.Println("Você não pode dividir por 0")
		}
	}

	_, err = dividir(-10, 10)
	if err != nil {
		fmt.Println("Err: ", err)
	}

}
