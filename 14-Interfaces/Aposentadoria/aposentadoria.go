package main

import "fmt"

type pessoas interface {
	aposentado() bool
}

type basePessoa struct {
	nome  string
	idade uint8
}

type homem struct {
	basePessoa
}

func (h homem) aposentado() bool {
	return h.idade >= 65
}

type mulher struct {
	basePessoa
}

func (m mulher) aposentado() bool {
	return m.idade >= 60
}

func verificaAposentadoria(p pessoas) {
	fmt.Println("Está aposentado(a)?", p.aposentado())
}

func main() {
	lucas := basePessoa{"Lucas", 30}
	pedro := basePessoa{"Pedro", 65}
	maria := basePessoa{"Maria", 30}
	camila := basePessoa{"Camila", 60}

	h := homem{lucas}
	h2 := homem{pedro}
	m := mulher{maria}
	m2 := mulher{camila}

	verificaAposentadoria(h)
	verificaAposentadoria(h2)
	verificaAposentadoria(m)
	verificaAposentadoria(m2)

}
