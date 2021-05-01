package main

import (
	"fmt"
	"math"
)

// Inteface forma que tem um metodo area que retorna um float64
// Assim todas as formas(retangulo e circulo) necessitam de
// um metodo area que não recebe parametros e devolve um float64
// para se encaixarem nessa interface
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
