package main

import (
	"fmt"
	"math"
)

type Operator func(x float64) float64

// Eu passo uma outra função para aqui dentro
func Map(op Operator, a []float64) []float64 {
	res := make([]float64, len(a))
	for i, x := range a {
		res[i] = op(x) //aqui eu chamo a função
	}
	return res
}

func main() {
	op := math.Abs //Função que retorna apenas absolutos
	a := []float64{-1, -2}
	b := Map(op, a)
	fmt.Println(b) // [1 2]

	c := Map(func(x float64) float64 { return 10 * x }, b)
	fmt.Println(c) // [10, 20]
}
