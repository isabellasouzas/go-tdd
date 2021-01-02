package contexto

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func (s *StubStore) Fetch() string {
	return s.response
}

type StubStore struct {
	response string
}

func TestHandler(t *testing.T) {
	data := "olá, mundo"
	srv := Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	srv.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`resultado "%s", esperado "%s"`, response.Body.String(), data)
	}

}
