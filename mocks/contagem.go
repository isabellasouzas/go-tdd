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

const ultimaPalavra = "Go!"
const inicioContagem = 3

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
