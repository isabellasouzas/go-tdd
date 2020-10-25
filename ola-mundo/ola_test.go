package main

import "testing"

func TestOla(t *testing.T) {
	resultado := ola("Chris")
	esperado := "Olá, Chris"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'.", resultado, esperado)
	}
}
