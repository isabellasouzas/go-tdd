package estruturas_metodos_e_interfaces

import "testing"

func TestPerimetro(t *testing.T) {
	resultado := Perimetro(10.0, 10.0)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
	}
}

func TestArea(t *testing.T) {
	retangulo := Retangulo{12.0, 6.0}
	resultado := Area(retangulo)
	esperado := 72.0

	if resultado != esperado {
		t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
	}
}
