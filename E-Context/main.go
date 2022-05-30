package main

import (
	"context"
	"fmt"
	"time"
)

/*
	Context serve para definir o tempo limite que as aplicações vão suportar

*/
func numeros(v chan<- int) {
	for i := 0; i < 10; i++ {
		v <- i
		fmt.Printf("Numero %d escrito do channel \n", i)
	}
	close(v)
}

func main() {
	ctx, cf := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 5) //Após 5 segundos o contexto é fechado
		cf()
		fmt.Println("Timeout!")
	}()

	c := make(chan int, 3)
	go numeros(c)

loopNum:
	for {
		select {
		case <-ctx.Done(): //Quando o contexto está cancelado
			break loopNum

		case v, ok := <-c:
			if ok {
				fmt.Printf("Numero %d lido do channel \n", v)
				time.Sleep(time.Second * 1)
			}
		}
	}
}
