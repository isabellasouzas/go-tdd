package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz ola para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris")
		esperado := "Olá, Chris"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("'Mundo' como padrao para uma 'string' vazia", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá, mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

}
