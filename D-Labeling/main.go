package main

import (
	"fmt"
	"time"
)

func main() {
	// um apelido para o loop
	// ATENÇÃO: Não é possivel saltar sobre variaveis
loopNum:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("Labeling -- j = %d i = %d\n", j, i)
			if j == 3 {
				fmt.Println("--- CONTINUE ---")
				continue loopNum
			}
			if i == 3 {
				fmt.Println("--- BREAK ---")
				break loopNum
			}

			if i == 5 {
				goto outraMensagem
			}
		}
	}

	goto mensagem
	fmt.Println("Isso é pulado") //É ELIMINADO AO COMPILAR

mensagem:
	fmt.Println("Apelido para o código")

outraMensagem:
	fmt.Println("Mas cuidado que ela não precisa ser chamada para executar")

	gotoNoSelect()
	gotoSwitch()
}

func gotoNoSelect() {
	timeout := time.After(3 * time.Second)

	messages := make(chan string)
	go func() {
		messages <- "ping1"
		messages <- "ping2"
		messages <- "ping3"
	}()

loop:
	select {
	case <-timeout:
		fmt.Println("timeout")

	case m := <-messages:
		fmt.Println(m)
		goto loop
	}
}

func gotoSwitch() {
	var i = 2
label:
	switch i {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
		i = 1
		goto label
	case 3:
		fmt.Println("case 3")
	case 4:
		fmt.Println("case 4")

	}
}
