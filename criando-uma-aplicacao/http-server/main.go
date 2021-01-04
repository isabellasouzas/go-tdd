package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(ServidorJogador)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("nao foi possivel escutar a porta 5000 '%v'", err)
	}
}
