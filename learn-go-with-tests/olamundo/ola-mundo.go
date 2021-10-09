package main

import (
	"fmt"
	e "go-basics/learn-go-with-tests/inteiros"
)

const spanish = "spanish"
const french = "french"
const prefixHelloInPortuguese = "Olá, "
const prefixHelloInSpanish = "Hola, "
const prefixHelloInFrench = "Ui, "

func Ola(name string, language string) string {
	if name == "" {
		name = "Mundo"
	}

	return prefixodeSaudacao(language) + name
}

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case french:
		prefixo = prefixHelloInFrench
	case spanish:
		prefixo = prefixHelloInSpanish
	default:
		prefixo = prefixHelloInPortuguese
	}
	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
	e.ExampleAdicionar()
}
