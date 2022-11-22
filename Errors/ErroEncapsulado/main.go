package main

import (
	"errors"
	"fmt"
)

type ProdutoErr struct {
	Produto string
	Codigo  int
}

func (pe ProdutoErr) Error() string {
	return fmt.Sprintf("Ocorreu um erro %s, %d", pe.Produto, pe.Codigo)
}

func fazAlgoComProduto() error {
	return ProdutoErr{"ProdutoX", 123}
}

func main() {
	err := fazAlgoComProduto()
	fmt.Println(err)

	var pe ProdutoErr
	if errors.As(err, &pe) { // Percorre a "pilha" que é do tipo da struct
		fmt.Println(pe.Produto)
		fmt.Println(pe.Codigo)
	}
}
