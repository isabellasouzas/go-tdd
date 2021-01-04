package main

import (
	"fmt"
	"net/http"
)

func ServidorJogador(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
