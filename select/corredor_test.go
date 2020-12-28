package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {

	servidorLento := criarServidorComAtraso(20 * time.Microsecond)
	servidorRapido := criarServidorComAtraso(0 * time.Microsecond)

	defer servidorLento.Close()
	defer servidorRapido.Close()

	URLLenta := servidorLento.URL
	URLRapida := servidorRapido.URL

	esperado := URLRapida
	resultado := Corredor(URLLenta, URLRapida)

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
	servidorLento.Close()
	servidorRapido.Close()
}

func criarServidorComAtraso(atraso time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(atraso)
		writer.WriteHeader(http.StatusOK)
	}))
}
