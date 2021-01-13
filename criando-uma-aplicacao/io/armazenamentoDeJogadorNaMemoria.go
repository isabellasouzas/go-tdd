package main

func NovoArmazenamentoDeJogadorNaMemoria() *ArmazenamentoDeJogadorNaMemoria {
	return &ArmazenamentoDeJogadorNaMemoria{map[string]int{}}
}

type ArmazenamentoDeJogadorNaMemoria struct {
	armazenamento map[string]int
}

func (i *ArmazenamentoDeJogadorNaMemoria) PegaLiga() []Jogador {
	var liga []Jogador
	for nome, vitorias := range i.armazenamento {
		liga = append(liga, Jogador{nome, vitorias})
	}
	return liga
}

func (i *ArmazenamentoDeJogadorNaMemoria) SalvaVitoria(nome string) {
	i.armazenamento[nome]++
}

func (i *ArmazenamentoDeJogadorNaMemoria) PegaPontuacaoDoJogador(nome string) int {
	return i.armazenamento[nome]
}
