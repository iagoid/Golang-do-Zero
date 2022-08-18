package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var waittime sync.WaitGroup
var atmCounter uint64
var counter uint64

// Variáveis Atômicas são utilizadas para controlar o estado.
// Além disso, também evita race condition que permitem que duas ou mais Goroutines acessem fontes idênticas.
// Os contadores atômicos estão disponíveis para várias goroutines.
// O principal meio para gerenciar estados na linguagem go é a comunicação por meio de canais.
func hike(S string) {
	for i := 1; i < 1000; i++ {
		atomic.AddUint64(&atmCounter, 1) // Sempre retorna 4995
		counter += 1                     // Pode variar o retorno
	}
	waittime.Done()
}

func main() {
	waittime.Add(5)

	go hike("cat: ")
	go hike("dog: ")
	go hike("duck: ")
	go hike("cow: ")
	go hike("horse: ")

	waittime.Wait()
	fmt.Println("Último valor atômico:", atmCounter)
	fmt.Println("Último valor do contador:", counter)
}
