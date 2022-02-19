package main

import "fmt"

type Document interface {
	Doc() string
}

type Pessoa struct {
	Nome   string
	Status bool
}

type PessoaFisica struct {
	Pessoa
	CPF string
}
type PessoaJuridica struct {
	Pessoa
	CNPJ string
}

// A função String faz com que isso seja printado toda vez que eu fizer print da interface
func (p Pessoa) String() string {
	return fmt.Sprint("Salve " + p.Nome)
}

func (p PessoaFisica) Doc() string {
	return fmt.Sprintf("Olá " + p.Nome + " O seu CPF é " + p.CPF)
}

func (p PessoaJuridica) Doc() string {
	return fmt.Sprintf("Olá " + p.Nome + " O seu CNPJ é " + p.CNPJ)
}

//Asserção de tipo
func show(d Document) {
	fmt.Println(d) // chama a função string
	fmt.Println(d.Doc())

	switch d.(type) {
	case PessoaFisica:
		fmt.Println(d.(PessoaFisica).Status)
	case PessoaJuridica:
		fmt.Println(d.(PessoaJuridica).Status)
	}

	// Ou
	// if p, ok := d.(PessoaFisica); ok { // O p não é mais interface, é do tipo pessoaFisica
	// 	fmt.Println("Seu status seu é", p.Status)
	// } else if p, ok := d.(PessoaJuridica); ok { // O p não é mais interface, é do tipo pessoaJuridica
	// 	fmt.Println("O status da empresa é", p.Status)
	// }
}

func main() {
	p1 := PessoaFisica{
		CPF:    "000.000.000-00",
		Pessoa: Pessoa{Nome: "Iago", Status: true}}
	show(p1)

	p2 := PessoaJuridica{
		CNPJ:   "111111111111",
		Pessoa: Pessoa{Nome: "Lydians", Status: true}}
	show(p2)
}
