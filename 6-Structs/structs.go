package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome = "Iago"
	u.idade = 19
	fmt.Println(u)

	// Adicionando um struct dentro do outro
	endereco := endereco{"Rua dos bobos", 123}

	usuario2 := usuario{"Matheus", 51, endereco}
	fmt.Println(usuario2)

	// Caso eu não tenha todos os dados
	usuario3 := usuario{idade: 32, endereco: endereco}
	fmt.Println(usuario3)

}
