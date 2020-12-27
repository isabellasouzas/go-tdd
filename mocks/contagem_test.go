package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestContagem(t *testing.T) {

	t.Run("imprime 3 até Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Contagem(buffer, &SpyContagemOperacoes{})

		resultado := buffer.String()
		esperado := `3
2
1
Go!`
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("pausa antes de cada impressão", func(t *testing.T) {
		spyImpressoeaSleep := &SpyContagemOperacoes{}
		Contagem(spyImpressoeaSleep, spyImpressoeaSleep)

		esperado := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(esperado, spyImpressoeaSleep.Chamadas) {
			t.Errorf("esperado %v, resultado %v", esperado, spyImpressoeaSleep.Chamadas)
		}
	})

}

func TestSleeperConfiguravel(t *testing.T) {
	tempoPausa := 5 * time.Second

	TempoSpy := &TempoSpy{}
	sleeper := SleeperConfiguravel{tempoPausa, TempoSpy.Pausa}
	sleeper.Pausa()

	if TempoSpy.duracaoPausa != tempoPausa {
		t.Errorf("deveria ter pausado por %v, mas pausou por %v", tempoPausa, TempoSpy.duracaoPausa)
	}
}
