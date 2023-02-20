package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	// Usa uma goroutine para verificar as modificações
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				// Monitora apenas os eventos de escrita
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("File modified:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()

	// Adiciona o arquivo que eu estou verificando
	err = watcher.Add("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
