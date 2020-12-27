package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Pausa()
}

//mock
type SleeperSpy struct {
	Chamadas int
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

type TempoSpy struct {
	duracaoPausa time.Duration
}

type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

const ultimaPalavra = "Go!"
const inicioContagem = 3
const escrita = "escrita"
const pausa = "pausa"

func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
}

func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}

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

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}

	sleeper.Pausa()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperConfiguravel{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}
