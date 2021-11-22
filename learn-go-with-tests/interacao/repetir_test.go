package interacao

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a")
	expect := "aaaaa"

	if repeticoes != expect {
		t.Errorf("esperado '%s' mas obteve '%s'", expect, repeticoes)
	}
}
