package main

import (
	"fmt"
	"net/http"
)

func ServidorJogador(w http.ResponseWriter, r *http.Request) {
	jogador := r.URL.Path[len("/jogadores/"):]

	if jogador == "Maria" {
		fmt.Fprint(w, "20")
		return
	}

	if jogador == "Pedro" {
		fmt.Fprint(w, "10")
	}
}
