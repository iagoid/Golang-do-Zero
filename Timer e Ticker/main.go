package main

import (
	"fmt"
	"time"
)

// Timers são para quando você quer fazer algo uma vez no futuro
// Tickers são para quando você quer fazer algo repetidamente em intervalos regulares.
func main() {
	timer()
	ticker()
}
func timer() {
	// Os temporizadores representam UM único evento no futuro.
	// Você informa ao cronômetro quanto tempo deseja esperar
	// e ele fornece um canal que será notificado nesse momento.
	// Este temporizador irá esperar 2 segundos.
	timer := time.NewTimer(time.Second * 2)
	go func() {
		<-timer.C // Espera o tempo passar
		fmt.Println("Timer fired")
	}()
	stopped := timer.Stop() // Para o cronometro
	if stopped {
		fmt.Println("Timer stopped")
	}
	time.Sleep(2 * time.Second)
}

func ticker() {
	//Os tickers são excepcionalmente úteis quando você precisa executar uma ação repetidamente
	// em determinados intervalos de tempo. Podemos usar tickers,
	// em combinação com goroutines, para executar essas tarefas em segundo plano em nossos aplicativos.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
