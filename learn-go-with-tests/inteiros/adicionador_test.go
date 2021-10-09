package inteiros

import (
	"testing"
)

func TestAdicianador(t *testing.T) {
	soma := Adicionar(2, 2)
	expect := 4

	if soma != expect {
		t.Errorf("esperado '%d', resultado '%d'", expect, soma)
	}
}
