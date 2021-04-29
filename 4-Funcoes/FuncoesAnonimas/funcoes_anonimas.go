package main

import "fmt"

func main() {
	// Declara e executa ela automaticamente

	func(texto string) {
		fmt.Println("Recebido ->", texto)
	}("Sem retorno")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parametro")

	fmt.Println(retorno)
}
