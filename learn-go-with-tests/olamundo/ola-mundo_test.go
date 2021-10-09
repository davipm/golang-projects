package main

import "testing"

func TestOla(t *testing.T) {
	verifyCorrectMessage := func(t *testing.T, result, expect string) {
		t.Helper()
		if result != expect {
			t.Errorf("resultado '%s', esperado '%s'", result, expect)
		}
	}

	t.Run("should say hello in spanish", func(t *testing.T) {
		result := Ola("Chris", "spanish")
		expect := "Hola, Chris"
		verifyCorrectMessage(t, result, expect)
	})

	t.Run("should say hello in french", func(t *testing.T) {
		result := Ola("Chris", "french")
		expect := "Ui, Chris"
		verifyCorrectMessage(t, result, expect)
	})

	t.Run("should say hello in portuguese", func(t *testing.T) {
		result := Ola("", "")
		expect := "Olá, Mundo"
		verifyCorrectMessage(t, result, expect)
	})
}
