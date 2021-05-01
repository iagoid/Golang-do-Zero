package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	// Coloca a quantidade de goroutine que fazem parte do grupo de espera
	waitGroup.Add(2)

	// O Done tira 1 do contador
	go func() {
		escrever("Olá mundo")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em go")
		waitGroup.Done() // -1
	}()

	// Espera a contagem chegar em 0
	waitGroup.Wait()
	escrever("As duas funcões anteriores foram executadas")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
