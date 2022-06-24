package main

import (
	"fmt"
)

func main() {
	const (
		Dollar = iota
		Euro
	)

	convertions := [2]float64{
		Euro:   5.53,
		Dollar: 5.24,
	}

	fmt.Print("R$1 vale $", convertions[Dollar])
	fmt.Print("\nR$1 vale €", convertions[Euro])

}
