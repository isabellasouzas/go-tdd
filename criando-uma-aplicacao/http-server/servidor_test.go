package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestObterJogador(t *testing.T) {
	t.Run("retornar resultado de Maria", func(t *testing.T) {
		requisicao, _ := http.NewRequest(http.MethodGet, "/jogadores/Maria", nil)
		resposta := httptest.NewRecorder()

		ServidorJogador(resposta, requisicao)

		recebido := resposta.Body.String()
		esperado := "20"

		if recebido != esperado {
			t.Errorf("recebido '%s', esperado '%s'", recebido, esperado)
		}

	})
	t.Run("retornar resultado Pedro", func(t *testing.T) {
		requisicao, _ := http.NewRequest(http.MethodGet, "/jogador/Pedro", nil)
		resposta := httptest.NewRecorder()

		ServidorJogador(resposta, requisicao)

		recebido := resposta.Body.String()
		esperado := "10"

		if recebido != esperado {
			t.Errorf("recebido '%s', esperado '%s'", recebido, esperado)
		}
	})
}
