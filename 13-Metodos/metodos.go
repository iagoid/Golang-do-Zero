package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// Criamos metodos para dar ações ao usuário
// Eles estão "grudados" ao tipo do usuario
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no BD\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// Referencia  a estrutura diretamente, sem criar uma cópia
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Iago", 19}
	usuario1.salvar()
	maiorIdadade1 := usuario1.maiorDeIdade()
	fmt.Println(maiorIdadade1)

	usuario2 := usuario{"Cláudio", 30}
	usuario2.salvar()
	maiorIdadade2 := usuario2.maiorDeIdade()
	fmt.Println(maiorIdadade2)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
