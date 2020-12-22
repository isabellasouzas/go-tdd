package estruturas_metodos_e_interfaces

import "testing"

func TestArea(t *testing.T) {
	testesArea := []struct {
		forma    Forma
		esperado float64
	}{
		{forma: Retangulo{Largura: 12, Altura: 6}, esperado: 72.0},
		{forma: Circulo{Raio: 10}, esperado: 314.1592653589793},
		{forma: Triangulo{Base: 12, Altura: 6}, esperado: 36.0},
	}
	for _, tt := range testesArea {
		resultado := tt.forma.Area()
		if resultado != tt.esperado {
			t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, resultado, tt.esperado)
		}
	}
}
