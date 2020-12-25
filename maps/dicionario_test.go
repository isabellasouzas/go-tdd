package maps

import "testing"

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "isso é apenas um teste"}

	resultado := Dicionario.Busca(dicionario, "teste")
	esperado := "isso é apenas um teste"

	comparaStrings(t, resultado, esperado)
}

func comparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", resultado, esperado, "teste")
	}
}
