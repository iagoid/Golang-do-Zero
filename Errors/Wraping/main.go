package main

import (
	"errors"
	"fmt"
	"os"
)

// Wrapping de erro
func verificaArquivo(nome string) error {
	f, err := os.Open(nome)
	if err != nil {
		return fmt.Errorf("erro ao abrir %w", err)
	}
	f.Close()
	fmt.Println("Arquivo Existe")
	return nil
}

func main() {
	err := verificaArquivo("naoExiste.txt")
	if errors.Is(err, os.ErrNotExist) { // se o erro encapsulado é o mesmo
		fmt.Println(err)
		fmt.Println(errors.Unwrap(err))
	}
}
