package ponteiros_erros

import "fmt"

type Bitcoin int

type Stringer interface {
	String() string
}

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Depositar(quantidade Bitcoin) {
	fmt.Printf("O endereço do saldo no Depositar é %v \n", &c.saldo)
	c.saldo += quantidade
}

func (c *Carteira) Saldo() int {
	return c.saldo
}

func (b Bitcoin) String() string {
	fmt.Sprintf("%d BTC", b)
}
