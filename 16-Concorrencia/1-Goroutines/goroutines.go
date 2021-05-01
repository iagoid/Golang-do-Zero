package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA É DIFERENTE DE PARALELISMO
	// Ao coloca go na frente da função ela fica
	// disputando com a outra função, sem esperar que
	// a primeira função termine para iniciar a outra
	go escrever("Olá mundo")
	escrever("Programando em go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
