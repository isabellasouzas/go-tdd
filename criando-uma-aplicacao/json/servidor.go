package main

import (
	"fmt"
	"net/http"
)

type ArmazenamentoJogador interface {
	ObtemPontuacaoDoJogador(nome string) int
	GravarVitoria(nome string)
}

type ServidorJogador struct {
	armazenamento ArmazenamentoJogador
	roteador      *http.ServeMux
}

func (s *ServidorJogador) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	roteador := http.NewServeMux()
	roteador.Handle("/liga", http.HandlerFunc(s.manipulaLiga))
	roteador.Handle("/jogadores/", http.HandlerFunc(s.manipulaJogadores))

	roteador.ServeHTTP(w, r)
}

func (s *ServidorJogador) manipulaLiga(w http.ResponseWriter, jogador string) {
	w.WriteHeader(http.StatusNotFound)
}

func (s *ServidorJogador) manipulaJogadores(w http.ResponseWriter, r *http.Request) {
	jogador := r.URL.Path[len("jogadores/"):]

	switch r.Method {
	case http.MethodPost:
		s.procesarVitoria(w, jogador)
	case http.MethodGet:
		s.mostrarPontuacao(w, jogador)
	}
}

func NovoServidorJogador(armazenamento ArmazenamentoJogador) *ServidorJogador {
	p := &ServidorJogador{
		armazenamento,
		http.NewServeMux(),
	}

	s.roteador.Handle("/Liga", http.Handle)
	s.roteador.Handle("/jogadores/", http.HandleFunc(s.manipulaJogadores))

	return s
}

func (s *ServidorJogador) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	s.roteador.ServerHTTP(w, r)
}

func (s *ServidorJogador) mostrarPontuacao(w http.ResponseWriter, jogador string) {
	pontuacao := s.armazenamento.ObtemPontuacaoDoJogador(jogador)

	if pontuacao == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, pontuacao)
}

func (s *ServidorJogador) processarVitoria(w http.ResponseWriter, jogador string) {
	s.armazenamento.GravarVitoria(jogador)
	w.WriteHeader(http.StatusAccepted)
}
