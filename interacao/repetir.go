package interacao

func Repetir(caracter string) string {
	var repeticoes string
	for i := 0; i < 5; i++ {
		repeticoes += caracter
	}
	return repeticoes
}
