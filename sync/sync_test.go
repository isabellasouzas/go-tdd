package sync

import (
	"testing"
)

func TestContador(t *testing.T) {
	t.Run("incrementar o contador 3 vezes resulta no valor 3", func(t *testing.T) {
		contador := Contador{}
		contador.Incrementa()
		contador.Incrementa()
		contador.Incrementa()

		if contador.Valor() != 3 {
			t.Errorf("resultado %d, esperado %d", contador.Valor(), 3)
		}
	})

	func VerificaContador(t*testing.T, resultado Contador, esperado	int) {
		t.Helper()
		if resultado.Valor() != esperado {
			t.Errorf("resultado %d, esperado %d", resultado.Valor(), esperado)
		}
	}
}