package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const TAMANHO = 10
	var values [TAMANHO]string

	cond := sync.NewCond(&sync.Mutex{})

	// Adiciona ao array uma contagem de 10 letras depois de esperar por um tempo aleatório
	for i := 0; i < TAMANHO; i++ {
		d := time.Second * time.Duration(rand.Intn(10)) / 10
		go func(i int) {
			time.Sleep(d)
			cond.L.Lock()
			values[i] = string('a' + i)
			cond.L.Unlock()
			cond.Signal() //Manda a notificação para o mais antigo, uma de cada vez
		}(i)
	}

	// Realiza a verificação se o array está totalmente preenchido
	checkCondition := func() bool {
		fmt.Println(values)
		for i := 0; i < TAMANHO; i++ {
			if values[i] == "" {
				return false
			}
		}
		return true
	}

	// Aguarda o programa finalizar quando checkCondition for verdadeiro
	cond.L.Lock()
	defer cond.L.Unlock()
	for !checkCondition() {
		cond.Wait()
	}
}
