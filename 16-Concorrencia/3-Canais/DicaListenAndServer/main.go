package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{}
}

func (s *Server) ListenAndServer() {
free:
	for {
		select {
		case <-s.quitch:
			fmt.Println("Saindo...")
			// usar apenas em casos onde eu tiver lógica abaixo do for
			// senão posso utilizar um return
			break free
		default:
		}
	}

	fmt.Println("Busca finlizada")
}

func (s *Server) quit() {
	close(s.quitch)
}

func main() {
	s := &Server{
		quitch: make(chan struct{}),
	}

	go func() {
		time.Sleep(2 * time.Second)
		s.quit()
	}()

	s.ListenAndServer()

}
