package main

import (
	"fmt"
	"math"
)

func main() {
	var peso uint8 = 255

	if peso <= math.MaxUint8 {
		fmt.Println("Seu peso é ", peso)
	} else {
		fmt.Println("Não é possível usar esse peso")
	}
}
