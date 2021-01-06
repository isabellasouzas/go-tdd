package main

func NovoArmazenamentoDeJogadorNaMemoria() *ArmazenamentoDeJogadorNaMemoria {
	return &NovoArmazenamentoDeJogadorNaMemoria{map[string]int{}}
}

type ArmazenamentoDeJogadorNaMemoria struct {
	armazenamento map[string]int
}

func (a *ArmazenamentoDeJogadorNaMemoria) GravarVitoria(nome string) {
	a.armazenamento[nome]++
}

func (a *ArmazenamentoDeJogadorNaMemoria) ObtemPontuacaDoJogador(nome string) {
	return a.armazenamento[nome]
}
