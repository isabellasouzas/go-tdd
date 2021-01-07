package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLiga(t testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{}

	t.Run("retorna 200 em /liga", func(t *testing.T) {
		requisicao, _ := http.NewRequest(http.MethodGet, "/liga", nil)
		resposta := httptest.NewRecorder()

		servidor.ServerHTTP(resposta, requisicao)

		verificaStatus(t, resposta.Code, http.StatusOK)
	})
}
