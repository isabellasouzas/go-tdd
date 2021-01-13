package main

import (
	"strings"
	"testing"
)

func TestSistemaDeArquivoDermazenamento(t *testing.T) {
	t.Run("/liga de um leitor", func(t *testing.T) {
		bancoDeDados := strings.NewReader(`[
				{"Nome": "Cleo", "Vitorias": 10},
				{"Nome": "Chris", "Vitorias": 33}]`)

		armazenamento := SistemasDeArquivoDeArmazenamentoDoJogador{bancoDeDados}

		recebido := armazenamento.PegaLiga()

		esperado := []Jogador{
			{"Cleo", 10},
			{"Chris", 33},
		}

		defineLiga(t, recebido, esperado)
	})
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) PegaLiga() []Jogador {
	liga, _ := NovaLiga(f.bancoDeDados)
	return liga
}
