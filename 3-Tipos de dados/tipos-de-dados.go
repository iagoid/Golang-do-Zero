package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64 de acordo com o numero de bits
	var numero1 int8 = 127
	var numero2 int16 = 32767
	var numero3 int32 = 2147483647
	var numero4 int64 = 9223372036854775807
	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)

	// Caso eu não especifique ele utiliza de acordo com a arquitetura do computador
	numero5 := 100000000000
	fmt.Println(numero5)

	// int sem sinal
	var numero6 uint16 = 65535
	fmt.Println(numero6)

	// alias
	// int32 = rune
	var numero7 rune = 70000
	fmt.Println(numero7)
	// int8 = rune
	var numero8 byte = 80
	fmt.Println(numero8)

	//float32, float64
	var real1 float32 = 3.4028234663852886e+38
	fmt.Println(real1)
	var real2 float64 = 1.7976931348623157e+308
	fmt.Println(real2)

	// char
	// O char é convertido para um numero da tabela ASCII
	char := 'B'
	fmt.Println(char)

	// boolean
	var booleano bool
	fmt.Println(booleano)

	// error
	var err error = errors.New("Erro interno")
	fmt.Println(err)

	casosEspecificaos()
}

// Repare que ao incrementar um numero após seu limite o contador volta a ser 0
func casosEspecificaos() {
	var numero uint8 = 255

	numero++
	fmt.Println(numero)
}
