package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá mundo", canal)
	fmt.Println("Depois da funçao escrever ser chamada")
	// Espera que  o canal receber um valor
	// Só vai para outra linha quando o valor é passado
	// LEMBRE-SE: passar e receber são ações bloqueantes
	mensagem := <-canal
	fmt.Println(mensagem)

	// Eu posso fazer um for para pegar todos os valores
	// Passados pelo canal, porém se eu utilizar um for
	// maior do que a quantidade retornada ocorre um deadlock
	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	// OU

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("----- Programa finalizado ------")

}

func escrever(texto string, canal chan string) {
	time.Sleep((time.Second) * 5)
	for i := 0; i < 5; i++ {
		// Manda o valor para dentro do canal
		fmt.Println("Dentro do for")
		canal <- fmt.Sprint(texto, i)
		// O canal não espera chegar em outro for, enviando a mensagem antes
		time.Sleep((time.Second) * 3)
	}
	// Indica que o canal não envia/recebe dados
	close(canal)
}
