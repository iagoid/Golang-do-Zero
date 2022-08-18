package tests

import "testing"

// Por padrão o nome do teste inicia com "Test" e a primeira após letra maiuscula
func TestIsOdd_return_odd(t *testing.T) {
	var value int64 = 3

	numberType := IsOdd(value)

	if numberType != ODD_KEYORD {
		t.Errorf("%s - expected %s, received %s",
			t.Name(),
			ODD_KEYORD,
			numberType)
	}
	return
}

func TestIsOdd_return_even(t *testing.T) {
	var value int64 = 2

	numberType := IsOdd(value)

	if numberType != EVEN_KEYORD {
		t.Errorf("%s - expected %s, received %s",
			t.Name(),
			EVEN_KEYORD,
			numberType)
	}
	return
}

// Para rodar só dar um go test
// Caso eu vá em run package test (no inicio desse arquivo no vscode
// e ir no arquivo odd.go ele marca onde o teste passou
// go test -coverprofile=c.out ./...   	<-- Gera o arquivo c.out
// go tool cover -html=c.out			<-- Mostra os resultados no html
