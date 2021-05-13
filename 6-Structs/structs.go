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

/////////////////// Ponteiros receptores de método ///////////////////
// A grande diferença é que, se você definir um método com um receptor do valor,
// você NÃO vai conseguir fazer alterações na instância daquele tipo no qual o método foi definido.
// Se quisermos alterar a instância da variável, precisaremos defini-los como um receptor pointer:
type Creature struct {
	Species string
}

func (c Creature) Alter() {
	c.Species = ""
}

func (c *Creature) AlterPointer() {
	c.Species = "Dinossauro"
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

	// Ponteiros
	var creature Creature = Creature{Species: "shark"}

	fmt.Printf("1) %+v\n", creature)
	creature.Alter()
	fmt.Printf("2) %+v\n", creature)
	creature.AlterPointer()
	fmt.Printf("3) %+v\n", creature)

}
