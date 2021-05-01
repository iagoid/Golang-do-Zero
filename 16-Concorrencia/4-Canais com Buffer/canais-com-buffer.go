package main

import (
	"fmt"
)

func main() {
	// O canal tem uma capacidade de 3
	// Ele só bloquia quando atinge a capacidade máxima
	canal := make(chan string, 3)

	canal <- "Primeira mensagem"
	canal <- "Outra mensagem"
	canal <- "Só ocorre deadlock se houverem 4 mensagens"

	mensagem := <-canal
	mensagem2 := <-canal
	mensagem3 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)
}
