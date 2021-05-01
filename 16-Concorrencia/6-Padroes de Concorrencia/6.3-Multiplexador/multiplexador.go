package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Possuo dois canais diferentes retornados em uma mesma variavel
func main() {
	canal := multiplexar(escrever("Olá mundo"), escrever("Programando em Go"))

	for i := 0; i < 20; i++ {
		fmt.Println(<-canal)
	}

}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor Recebido: %s", texto)
			// Sorteio um numero pseudoaleatorio para verificar a diferença na execução
			time.Sleep((time.Millisecond) * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}

func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			// Verifica qual canal possui uma mensagem para ser lida
			select {
			case mensagem := <-canalEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()
	return canalDeSaida
}
