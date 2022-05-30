package main

import (
	"context"
	"fmt"
	"time"
)

// https://www.youtube.com/watch?v=mLISc6OGyFk

func main() {

	ctx := context.Background()            //contexto pai
	ctx, cancel := context.WithCancel(ctx) //gera um sinal de cancelamento
	// ctx, cancel := context.WithTimeout(ctx, time.Second*6) // gera um cancel que se cancela automatimamente após determinado tempo
	defer cancel()

	go func() { // em 6 segundos ele cancela o contexto (mudar os valores para testar)
		time.Sleep(time.Second * 6)
		cancel()
	}()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done(): //Quando o contexto está cancelado
		fmt.Println("Tempo exedido para bookar o quarto")
	case <-time.After(time.Second * 5): //Em 5 segundos o quarto é reservado
		fmt.Println("Quarto reservado com sucesso")
	}
}
