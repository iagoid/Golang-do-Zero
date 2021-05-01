package main

import "fmt"

// Init executa antes do main
// Posso ter uma função init por arquivo (não por pacote igual a main)
var n int

func init() {
	fmt.Println("Função Init")
	n = 12345
}

func main() {
	fmt.Println("Função Main")
	fmt.Println(n)
}
