package main

import (
	"fmt"
	"time"
)

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor Recebido: %s", texto)
			time.Sleep((time.Millisecond) * 500)
		}
	}()

	return canal
}

func main() {
	// Tenho infinitos canais criados
	// Serve para esconder a complexidade do código
	canal := escrever("Olá mundo!")

	for {
		fmt.Println(<-canal)
	}

}
