package main

import "fmt"

// Deslocamento de bits

// Constantes numéricas são valores de alta-precisão.
// Uma constante sem tipo pega o tipo necessário pelo seu contexto.
// Tente mostrar needInt(Big) também.

const (
	// Cria um grande número deslocando 1 bit para a esquerda 100 casas.
	// Em outras palavras, o número binário que é 1 seguido por 100 zeros.
	Big = 1 << 100
	// Desloque para a direita novamente 99 casas, então terminamos com 1 << 1 ou 2.
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	// fmt.Println(needInt(Big)) // Não funciona pois um int pode armazenar no máximo um inteiro de 64 bits
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	fmt.Println("=====================================================================")
	valor2 := 1 << 1         // 2
	valor4 := valor2 << 1    // 4
	valor8 := valor4 << 1    // 8
	valor128 := valor8 << 4  // 128
	valor64 := valor128 >> 1 // 64
	valor32 := valor64 >> 1  // 32
	valor16 := valor32 >> 1  // 16

	valor256 := valor16 << 4 // 256
	valor1 := valor256 >> 8  //1

	fmt.Println(valor1)
	fmt.Println(valor2)
	fmt.Println(valor4)
	fmt.Println(valor8)
	fmt.Println(valor16)
	fmt.Println(valor32)
	fmt.Println(valor64)
	fmt.Println(valor128)
	fmt.Println(valor256)

	fmt.Println("=====================================================================")

	valor24 := 3 << 3 // 3 -> 6 -> 12 -> 24
	fmt.Println(valor24)

	////////////////////////////////////////////////////////////////////////////////////////
	fmt.Println("=====================================================================")

	var x uint = 9  //0000 1001
	var y uint = 65 //0100 0001
	//  &     0000 0001 = 1
	//  |     0100 1001 = 73
	//  ^     0101 1000 = 72
	var z uint

	z = x & y
	fmt.Println("x & y  =", z)

	z = x | y
	fmt.Println("x | y  =", z)

	z = x ^ y
	fmt.Println("x ^ y  =", z)

	z = x << 1
	fmt.Println("x << 1 =", z)

	z = x >> 1
	fmt.Println("x >> 1 =", z)

}
