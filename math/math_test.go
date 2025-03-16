package math

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %2.f. Esperado: %2.f", total, 30.0)
	}
}
