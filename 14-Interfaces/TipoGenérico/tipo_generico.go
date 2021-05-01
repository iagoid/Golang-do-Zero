package main

import "fmt"

// Função com interface generica
// Posso passar qualquer tipo de dado
// CUIDADO: Não use em todos os casos senão vira GAMBIARRA
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	// OLHA A BAGUNÇA
	// PERDEU TOTALMENTE A SEGURANÇA DE TIPOS DO GO
	mapa := map[interface{}]interface{}{
		1:              "String",
		float32(12.55): true,
		true:           "Cleito",
		"Olá":          12,
	}
	fmt.Println(mapa)
}
