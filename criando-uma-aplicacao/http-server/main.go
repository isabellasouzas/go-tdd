package main

import (
	"log"
	"net/http"
)

type ArmazenamentoJogadorEmMemoria struct {
}

func (a *ArmazenamentoJogadorEmMemoria) ObterPontuacaoJogador(nome string) int {
	return 123
}

func main() {
	server := &ServidorJogador{&ArmazenamentoJogadorEmMemoria{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("nao foi possivel escutar a porta 5000 '%v'", err)
	}
}
