package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     uint
}

func (p *Pessoa) SetNome(nome string) *Pessoa {
	p.nome = nome
	return p
}

func (p *Pessoa) SetSobrenome(sobrenome string) *Pessoa {
	p.sobrenome = sobrenome
	return p
}

func (p *Pessoa) SetIdade(idade uint) *Pessoa {
	p.idade = idade
	return p
}

func (p *Pessoa) Print() {
	fmt.Printf("Olá, meu nome é %s %s e eu tenho %d anos.", p.nome, p.sobrenome, p.idade)
}

/* https://aprendagolang.com.br/2022/05/27/como-fazer-encadeamento-de-metodos-chaining/#respond
Embora essa técnica possa deixar nosso código mais limpo, ela também pode trazer um grande problema que é o tratamento de erros.
Quando um erro ocorre em outra liguagens, uma excessão é lançada e a sequencia de chamadas é quebrada,
ou seja, a execução do chaining para.

Como não temos excessões em Go, os métodos só podem retornar o ponteiro da struct,
evitando assim que possamos fazer um tratamento de erro adequado.

Utilize essa técnica somente em métodos simples, locais onde não existe a possibilidade de erro,
--> Métodos do tipo Set ou preparatórios.
*/

func main() {
	p := Pessoa{}
	p.SetNome("Iago").
		SetSobrenome("Dalmolin").
		SetIdade(20).
		Print()
}
