package enderecos

//Teste em outra pacote
// Faz o import, o . no começo indica que aquele pacote é chamado por padrão
import (
	// . "introducao-testes/enderecos"
	"testing"
)

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

//Começa com Test e NomeDaFunção
func TestTipoDeEndereco(t *testing.T) {
	// Rodar testes em paralelo
	// Ele roda em paralelo só quem tem essa chamada
	t.Parallel()

	cenarioTeste := []cenarioTeste{
		{"Rua ABC", "Rua"},
		{"rUa ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"aVeNiDa Paulista", "Avenida"},
		{"Estrada da Serra", "Estrada"},
		{"eStRaDa da Serra", "Estrada"},
		{"Rodovia 23364", "Rodovia"},
		{"rOdOvIa 123", "Rodovia"},
		{"Bairro da Fé", "Tipo Inválido"},
		{"", "Tipo Inválido"},
		{"RuaAvenidaEstradaRodovia", "Tipo Inválido"},
	}

	for _, cenario := range cenarioTeste {
		tipoEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("Endereço recebido diferente do esperado. \nRecebido: %s \nEsperado: %s",
				tipoEnderecoRecebido, cenario.retornoEsperado)
		}
	}

}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou")
	}
}
