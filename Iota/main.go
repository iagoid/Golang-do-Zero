package main

import "fmt"

// iota é um enumerador de tipos
const (
	A = iota + 1 //1
	_            //ignorado
	C            // 3
	D            // 4
)

type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1 << 0 que é 00000001
	IgChocolate                         // 1 << 1 que é 00000010
	IgNuts                              // 1 << 2 que é 00000100
	IgStrawberries                      // 1 << 3 que é 00001000
	IgShellfish                         // 1 << 4 que é 00010000
)

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

// Utilizando 2 ou mais iotas, ele incrementa os valores em colunas
const (
	Apple, Banana     = iota + 1, iota + 2 // Apple: 1, Banana: 2
	Cherimoya, Durian                      // Cherimoya: 2, Durian: 3
	Elderberry, Fig                        // Elderberry: 3, Fig: 4
)

func main() {

	fmt.Println(A, C, D) // "1 3 4"

	fmt.Println(IgEggs | IgChocolate | IgShellfish) //c Soma todos os Bytes --> 19 00010011

	fmt.Println(YB, "Bytes")

	direction()
}

// Passando uma string pela iota
// Passa o numero como uma string(como no Print)
type Direction int

const (
	North Direction = iota //0
	East                   //1
	South                  //2
	West                   //3
)

func direction() {
	var d Direction = North
	switch d {
	case North:
		fmt.Println(d, " goes up.")
	case South:
		fmt.Println(d, " goes down.")
	default:
		fmt.Println(d, " stays put.")
	}
}

// apenas tenha cuidado que o que é devolvido aqui precisa estar na mesma ordem que a constante
func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}
