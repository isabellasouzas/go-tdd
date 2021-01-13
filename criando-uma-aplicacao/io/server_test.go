package main

import (
	"io"
	"io/ioutil"
	"os"
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

		//read again
		recebido = armazenamento.PegaLiga()
		defineLiga(t, recebido, esperado)
	})
	t.Run("pegar pontuação do jogador", func(t *testing.T ){
		bancoDeDados := strings.NewReader( `[
					{"Nome": "Cleo", "Vitorias": 10},
					{"Nome": "Chris", "Vitorias": 33}]`)
		})
	armazenamento := SistemaDeArquivoDeArmazenamentoDoJogador{bancoDeDados}

	recebido := armazenamento.("Chris")
	esperado := 33

	definePontuacaoIgual(t, recebido, esperado)

}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) PegaLiga() []Jogador {
	liga, _ := NovaLiga(f.bancoDeDados)
	return liga
}

func criaArquivoTemporario(t *testing.T, dadoInicial string) (io.ReadWriteSeeker,  func())  {
	t.Helper()

	arquivotmp, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("não foi possivel escrever o arquivo temporário %v", err)
	}

	arquivotmp.Write([]byte(dadoInicial))

	removeArquivo := func() {
		arquivotmp.Close()
		os.Remove(arquivotmp.Name())
	}

	return arquivotmp, removeArquivo
}
