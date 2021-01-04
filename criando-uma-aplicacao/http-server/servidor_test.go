package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type EsbocoArmazenamentoJogador struct {
	pontuacoes        map[string]int
	registrosVitorias []string
}

func (e *EsbocoArmazenamentoJogador) ObterPontuacaoJogador(nome string) int {
	pontuacao := e.pontuacoes[nome]
	return pontuacao
}

func TestObterJogador(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{
			"Maria": 20,
			"Pedro": 10,
		},
		nil,
	}

	servidor := &ServidorJogador{&armazenamento}

	t.Run("retornar resultado de Maria", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacaoGet("Maria")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
		verificarCorpoRequisicao(t, resposta.Body.String(), "20")
	})
	t.Run("retornar resultado Pedro", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacaoGet("Pedro")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)
		verificarCorpoRequisicao(t, resposta.Body.String(), "10")
	})
	t.Run("retorna 404 para jogador não encontrado", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacaoGet("jorge")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		recebido := resposta.Code
		esperado := http.StatusNotFound

		if recebido != esperado {
			t.Errorf("recebido status '%d', esperado '%d'", recebido, esperado)
		}
	})
}

func TestArmazenamentoVitorias(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{},
		nil,
	}
	servidor := &ServidorJogador{&armazenamento}

	t.Run("retorna status 'aceito' para chamadas ao metodo POST", func(t *testing.T) {
		requisicao, _ := http.NewRequest(http.MethodPost, "/jogadores/Maria", nil)
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificarRespostaCodigoStatus(t, resposta.Code, http.StatusAccepted)
	})

}

func novaRequisicaoObterPontuacaoPost(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func novaRequisicaoObterPontuacaoGet(nome string) *http.Request {
	requisicao, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return requisicao
}

func verificarCorpoRequisicao(t *testing.T, recebido, esperado string) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("corpo da requisicao é invalido, obtive '%s' esperava %s", recebido, esperado)
	}
}

func verificarRespostaCodigoStatus(t *testing.T, recebido, esperado int) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("não recebeu o codigo de status http esperado, recebido %d, esperado %d", recebido, esperado)
	}
}
