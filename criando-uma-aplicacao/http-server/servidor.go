package main

import (
	"fmt"
	"net/http"
)

type AmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) int
	RegistrarVitoria(nome string)
}

type ServidorJogador struct {
	armazenamento AmazenamentoJogador
}

type ArmazenamentoJogadorEmMemoria struct {
	armazenamento map[string]int
}

func (s *ServidorJogador) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	jogador := request.URL.Path[len("/jogadores/"):]

	switch request.Method {
	case http.MethodPost:
		s.registrarVitoria(writer, jogador)
	case http.MethodGet:
		s.mostrarPontuacao(writer, jogador)
	}
}

func (s *ServidorJogador) mostrarPontuacao(writer http.ResponseWriter, jogador string) {
	pontuacao := s.armazenamento.ObterPontuacaoJogador(jogador)

	if pontuacao == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(writer, jogador)
}

func (s *ServidorJogador) registrarVitoria(w http.ResponseWriter, jogador string) {
	s.armazenamento.RegistrarVitoria(jogador)
	w.WriteHeader(http.StatusAccepted)
}
