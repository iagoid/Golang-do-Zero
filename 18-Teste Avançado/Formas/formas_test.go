package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	// Utilizado quando eu quero separar os testes, pois temos duas funções
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()
		if areaEsperada != areaRecebida {
			// Diferente do Errorf ele para o teste por aqui
			t.Fatalf("Área recebida diferente da esperada. \nRecebida: %f \nEsperada: %f",
				areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()
		if areaEsperada != areaRecebida {
			t.Fatalf("Área recebida diferente da esperada. \nRecebida: %f \nEsperada: %f",
				areaRecebida, areaEsperada)
		}
	})
}
