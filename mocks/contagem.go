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

const ultimaPalavra = "Go!"
const inicioContagem = 3

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

func Contagem(saida io.Writer) {
	for i := inicioContagem; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(saida, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprintln(saida, ultimaPalavra)
}

func main() {
	Contagem(os.Stdout)
}
