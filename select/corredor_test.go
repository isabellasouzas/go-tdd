package _select

import "testing"

func TestCorredor(t *testing.T) {
	URLLenta := "http://www.facebook.com"
	URLRapida := "http://www.quii.co.uk"

	esperado := URLRapida
	resultado := Corredor(URLLenta, URLRapida)

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
