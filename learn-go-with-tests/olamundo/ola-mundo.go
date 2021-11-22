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

/**
 * Hamming Distance: https://en.wikipedia.org/wiki/Hamming_distance
 *
 *
 * Hamming distance is a metric for comparing two binary data strings.
 *
 * While comparing two binary strings of equal length, Hamming distance
 * is the number of bit positions in which the two bits are different.
 * The Hamming distance between two strings, a and b is denoted as d(a,b)
 */

/**
 * @param {string} a
 * @param {string} b
 * @return {number}
 */
func hammingDistance(a, b string) int {
	if len(a) != len(b) {
		fmt.Println("Strings must be of the same length")
	}

	distance := 0

	for i := range a {
		if a[i] != b[i] {
			distance += 1
		}
	}

	return distance
}

func main() {
	fmt.Println(Ola("mundo", ""))
	e.ExampleAdicionar()

	fmt.Println(hammingDistance("david", "devod"))
}
