package main

import "testing"

func TestOla(t *testing.T) {
	t.Run("shold say hello", func(t *testing.T) {
		result := Ola("Chris")
		expect := "Ola, Chris"

		if result != expect {
			t.Errorf("resultado '%s', esperado '%s'", result, expect)
		}
	})

	t.Run("", func(t *testing.T) {

	})
}
