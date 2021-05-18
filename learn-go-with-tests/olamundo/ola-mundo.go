package main

import "fmt"

const prefixHelloInPortuguese = "Ola, "

func Ola(name string) string {
	return prefixHelloInPortuguese + name
}

func main() {
	fmt.Println(Ola("mundo"))
}
