package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = false

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	m := sync.Mutex{}
	c := sync.NewCond(&m)
	go func() {
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine1 wait")
			c.Wait() // Aguarda até receber um sinal
		}
		fmt.Println("goroutine1", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	go func() {
		c.L.Lock()
		for sharedRsc == false {
			fmt.Println("goroutine2 wait")
			c.Wait() // Aguarda até receber um sinal
		}
		fmt.Println("goroutine2", sharedRsc)
		c.L.Unlock()
		wg.Done()
	}()

	time.Sleep(2 * time.Second)

	c.L.Lock()
	fmt.Println("main goroutine ready")
	sharedRsc = true
	// Manda o sinal para todas as rotinas que algo aconteceu
	// Manda do mais recente, para o mais antigo
	// Ele manda ao "mesmo" tempo
	c.Broadcast()

	fmt.Println("main goroutine broadcast")
	c.L.Unlock()

	wg.Wait()
}
