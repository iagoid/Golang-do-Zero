package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal1"
		}
	}()

	go func() {
		for {
			time.Sleep((time.Second) * 2)
			canal2 <- "Canal2"
		}
	}()

	/*
		Com um for normal
		Ele executa a mensagem1 e após 2 segundos a mensagem2
		então a mensagem1 rapidamente e espera 2s para a mensagem2...
		O dever do select é fazer com que a mensagem1 possa ser executada
		mais vezes --> 4 mensagem1 para cada mensagem2
	*/
	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}

		// mensagem1 := <-canal1
		// fmt.Println(mensagem1)

		// mensagem2 := <-canal2
		// fmt.Println(mensagem2)
	}

}
