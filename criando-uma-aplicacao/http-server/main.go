package main

import (
	"log"
	"net/http"
)

func (a *ArmazenamentoJogadorEmMemoria) ObterPontuacaoJogador(nome string) int {
	return 123
}

func main() {
	servidor := &ServidorJogador{NovoArmazenamentoJogadorEmMemoria()}

	if err := http.ListenAndServe(":5000", servidor); err != nil {
		log.Fatalf("nao foi possivel escutar a porta 5000 '%v'", err)
	}
}
