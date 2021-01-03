package contexto

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func (s *StubStore) Fetch() string {
	return s.response
}

type StubStore struct {
	response string
}

func TestHandler(t *testing.T) {
	data := "olá, mundo"

	t.Run("retorna dados da store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`recebi "%s", quero "%s"`, response.Body.String(), data)

		}
	})
	t.Run("avisa a store para cancelar o trabalho se a requisiçãp for cancelada", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		store.assertWasCancelled()
	})
	t.Run("retorna dados da store", func(t *testing.T) {
		store := SpyStore{response: data}
		srv := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("resultado '%s', esperado '%s'", response.Body.String(), data)
		}

		if store.Cancelled {
			t.Error("não deveria ter cancelado a store")
		}
	})
	t.Run("avisa a store para cancelar o trabalho se a requisição for cancelada", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		srv.ServeHTTP(response, request)

		if response.written {
			t.Error("uma resposta não deveria ter sido escrita")
		}
	})
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Errorf("store não foi avisada para cancelar")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Errorf("store foi avisada para cancelar")
	}

}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("não implementado")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
