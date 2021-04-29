package main

import "fmt"

// NÃO REOMENDADO PARA FUNÇÕES QUE ENVOLVAM MUITOS NUMEROS
// POIS PODE OCORRER UM ESTOURO DE PILHA
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

/* BASICAMENTE:
fatorial(6)
 fatorial 5 * 6
 (fatorial 4 * 5) * 6
 ((fatorial 3 * 4) * 5) * 6
 (((fatorial 2 * 3) * 4) * 5) * 6
 ((((fatorial 1 * 2) * 3) * 4) * 5) * 6
 (((((fatorial 0 * 1) * 2) * 3) * 4) * 5) * 6
 (((((1 * 1) * 2) * 3) * 4) * 5) * 6
 ((((1 * 2) * 3) * 4) * 5) * 6
 (((2 * 3) * 4) * 5) * 6
 ((6 * 4) * 5) * 6
 (24 * 5) * 6
 120 * 6
*/

func fatorial(numero uint) uint {
	if numero == 0 {
		return 1
	}
	return fatorial(numero-1) * numero
}

func main() {
	numero := fibonacci(5)
	fmt.Println(numero)

	fatoriaResultado := fatorial(6)
	fmt.Println(fatoriaResultado)

}
