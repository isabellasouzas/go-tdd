package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type ArmazenamentoJogadorEmMemoria struct {
	armazenamento map[string]int
}

func TestRegistrarVitoriasEBuscarEstasVitorias(t *testing.T) {
	armazenamento := NovoArmazenamentoJogadorEmMemoria()
	servidor := ServidorJogador{armazenamento}
	jogador := "Maria"

	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoPontuacaoPost(jogador))
	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoPontuacaoPost(jogador))
	servidor.ServeHTTP(httptest.NewRecorder(), novaRequisicaoPontuacaoPost(jogador))

	resposta := httptest.NewRecorder()
	servidor.ServeHTTP(resposta, novaRequisicaoPontuacaoGet(jogador))
	verificarRespostaCodigoStatus(t, resposta.Code, http.StatusOK)

	verificarCorpoRequisicao(t, resposta.Body.String(), "3")
}

func NovoArmazenamentoJogadorEmMemoria() *ArmazenamentoJogadorEmMemoria {
	return &ArmazenamentoJogadorEmMemoria{map[string]int{}}
}

func (a *ArmazenamentoJogadorEmMemoria) RegistrarVitoria(nome string) {
	a.armazenamento[nome]++
}

func (a *ArmazenamentoJogadorEmMemoria) ObterPontuacaoJogador(nome string) int {
	return a.armazenamento[nome]
}
