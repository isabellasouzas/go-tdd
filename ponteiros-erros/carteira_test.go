package ponteiros_erros

import (
	"fmt"
	"testing"
)

func TestCarteira(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(Bitcoin(10))

	resultado := carteira.Saldo()

	fmt.Printf("O endereço do saldo no teste é %v \n", &carteira.saldo)
	esperado := Bitcoin(10)

	if resultado != esperado {
		t.Errorf("resultado %d, esperado %d", resultado, esperado)
	}
}
