package estruturas_metodos_e_interfaces

import "testing"

func TestArea(t *testing.T) {
	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
		t.Helper()
		resultado := forma.Area()

		if resultado != esperado {
			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
		}
	}

	t.Run("retângulos", func(t *testing.T) {
		retangulo := Retangulo{12, 6.0}
		verificaArea(t, retangulo, 72.0)
	})

	t.Run("círculos", func(t *testing.T) {
		circulo := Circulo{10}
		verificaArea(t, circulo, 314.1592653589793)
	})
}
