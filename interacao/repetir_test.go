package interacao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a")
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a")
	}
}

func ExampleRepetir() {
	repeticao := Repetir("a")
	fmt.Println(repeticao)
	// Output: aaaaa
}
