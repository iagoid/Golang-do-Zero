package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Envia uma requisição para o local host que se não for devolvida em 10 segundos cancela
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8000", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error send request: %v", err)
	}

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body) //resultado da request pro stdout
}
