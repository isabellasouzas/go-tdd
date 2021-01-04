package main

import (
	"fmt"
	"net/http"
)

type AmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) int
}

type ServidorJogador struct {
	armazenamento AmazenamentoJogador
}

func (s *ServidorJogador) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	jogador := request.URL.Path[len("/jogadores/"):]

	pontuacao := s.armazenamento.ObterPontuacaoJogador(jogador)

	if pontuacao == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(writer, pontuacao)
}

func (s *ServidorJogador) registrarVitoria(writer http.ResponseWriter, request *http.Request) {
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
