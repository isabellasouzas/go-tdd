package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GuardaJogador armazena informações sobre os jogadores
type GuardaJogador interface {
	PegaPontuacaoDoJogador(nome string) int
	SalvaVitoria(nome string)
	PegaLiga() []Jogador
}

// Jogador guarda o nome com o número de vitorias
type Jogador struct {
	Nome     string
	Vitorias int
}

// ServidorDoJogador é uma interface HTTP para informações dos jogadores
type ServidorDoJogador struct {
	armazenamento GuardaJogador
	http.Handler
}

const jsonContentType = "application/json"

// NovoServidorDoJogador cria um ServidorDoJogador com roteamento configurado
func NovoServidorDoJogador(armazenamento GuardaJogador) *ServidorDoJogador {
	p := new(ServidorDoJogador)

	p.armazenamento = armazenamento

	roteador := http.NewServeMux()
	roteador.Handle("/liga", http.HandlerFunc(p.ManipulaLiga))
	roteador.Handle("/jogadores/", http.HandlerFunc(p.ManipulaJogador))

	p.Handler = roteador

	return p
}

func (p *ServidorDoJogador) ManipulaLiga(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(p.armazenamento.PegaLiga())
	w.Header().Set("content-type", jsonContentType)
	w.WriteHeader(http.StatusOK)
}

func (p *ServidorDoJogador) ManipulaJogador(w http.ResponseWriter, r *http.Request) {
	jogador := r.URL.Path[len("/jogadores/"):]

	switch r.Method {
	case http.MethodPost:
		p.processaVitoria(w, jogador)
	case http.MethodGet:
		p.mostraPontuacao(w, jogador)
	}
}

func (p *ServidorDoJogador) mostraPontuacao(w http.ResponseWriter, jogador string) {
	pontuacao := p.armazenamento.PegaPontuacaoDoJogador(jogador)

	if pontuacao == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, pontuacao)
}

func (p *ServidorDoJogador) processaVitoria(w http.ResponseWriter, jogador string) {
	p.armazenamento.salvaVitorias(jogador)
	w.WriteHeader(http.StatusAccepted)
}

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	bancoDeDados io.ReadWriteSeeker
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) PegaLiga() []Jogador {
	f.bancoDeDados.Seek(0,0)
	liga, _ := NovaLiga(f.bancoDeDados)
	return liga
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) (nome string) int  {
	jogador := f.PegaLiga().Find(nome)

	if  jogador != nil {
		return  jogador.Vitorias
	}

	return 0
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) SalvaVitoria(nome string) {
	liga := f.PegaLiga()
	jogador :=liga.Find(nome)

	if  jogador != nil {
		jogador.Vitorias++
	}

	f.bancoDeDados.Seek(0, 0)
	json.NewEncoder(f.bancoDeDados).Encode(liga)
}