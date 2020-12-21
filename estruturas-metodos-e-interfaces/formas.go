package estruturas_metodos_e_interfaces

import "math"

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	Raio float64
}

func (r Retangulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}
