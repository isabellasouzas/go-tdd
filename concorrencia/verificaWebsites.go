package concorrencia

type VerificadorWebsites func(string) bool

func VerificaWebsites(vw VerificadorWebsites, urls []string) map[string]bool {
	resultados := make(map[string]bool)

	for _, url := range urls {
		resultados[url] = vw(url)
	}
	return resultados
}
