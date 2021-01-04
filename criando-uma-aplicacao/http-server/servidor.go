package main

import (
	"fmt"
	"net/http"
)

func ServidorJogador(w http.ResponseWriter, r *http.Request) {
	jogador := r.URL.Path[len("/jogadores/"):]

	fmt.Fprint(w, ObterPontuacaoJogador(jogador))

}

func ObterPontuacaoJogador(nome string) string {
	if nome == "Maria" {
		return "20"
	}

	if nome == "Pedro" {
		return "10"
	}

	return ""
}
