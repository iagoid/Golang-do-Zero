package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo %v ", err)
	}
	defer file.Close()

	bytesRead := make([]byte, 20)
	n, err := file.Read(bytesRead)
	if err != nil {
		log.Fatalf("Erro ao ler arquivo %v", err)
	}

	fmt.Printf("Foram lidos %d bytes: %s\n", n, string(bytesRead))

	// SEGUNDA LEITURA
	n, err = file.Read(bytesRead)
	if err != nil {
		if err == io.EOF {
			log.Printf("Leitura do arquivo completada. Total lido %d bytes\n", n)
		} else {
			log.Fatalf("Erro ao ler arquivo %v", err)
		}
	}

	fmt.Printf("Foram lidos %d bytes: %s\n", n, string(bytesRead))     // Não zera o array
	fmt.Printf("Foram lidos %d bytes: %s\n", n, string(bytesRead[:n])) // Pega só até onde foi lido

	// Leitura através de string
	reader := strings.NewReader("Água mole e pedra dura")
	n, _ = reader.Read(bytesRead)
	fmt.Printf("Foram lidos %d bytes: %s\n", n, string(bytesRead))

	leSite()

}

func leSite() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:80", "globo.com.br"))
	if err != nil {
		log.Fatalf("Erro ao ler site %v", err)
	}
	buf := make([]byte, 512)
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}
