package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {

	servidorLento :=httptest.NewServer(http.HandlerFunc(w http.ResponseWriter, r *http.Request){
		time.Sleep( 20 * time.Microsecond)
		w.WriteHeader(http.StatusOK)
	})

	servidorRapido := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
	}))

	URLLenta := servidorLento
	URLRapida := servidorRapido

	esperado := URLRapida
	resultado := Corredor(URLLenta, URLRapida)


	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
	servidorLento.Close()
	servidorRapido.Close()
}
