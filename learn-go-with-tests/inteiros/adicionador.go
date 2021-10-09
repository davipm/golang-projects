package inteiros

import "fmt"

func Adicionar(x, y int) int {
	return x + y
}

func ExampleAdicionar() {
	soma := Adicionar(1, 5)
	fmt.Println(soma)
}
