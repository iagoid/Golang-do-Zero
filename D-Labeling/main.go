package main

import "fmt"

func main() {
	// um apelido para o loop
loopNum:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("Labeling -- j = %d i = %d\n", j, i)
			if j == 3 && i == 1 {
				fmt.Println("===== Quebra loop =====")
				break loopNum
			}

			if i == 3 {
				goto outraMensagem
			}
		}
	}

	goto mensagem
	fmt.Println("Isso é pulado")

mensagem:
	fmt.Println("Apelido para o código")

outraMensagem:
	fmt.Println("Mas cuidado que ela não precisa ser chamada para executar")
}
