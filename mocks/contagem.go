package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

//mock
type SleeperSpy struct {
	Chamadas int
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

type SleeperPadrao struct {
}

const ultimaPalavra = "Go!"
const inicioContagem = 3
const escrita = "escrita"
const pausa = "pausa"

func (s *SpyContagemOperacoes) Pausa() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}

	for i := inicioContagem; i > 0; i-- {
		fmt.Fprintln(saida, i)
	}

	sleeper.Pausa()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper)
}
