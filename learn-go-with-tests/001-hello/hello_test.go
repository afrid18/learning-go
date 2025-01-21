package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to asifa", func(t *testing.T) {
		got := Hello("Asifa", "")
		want := "Hello, Asifa"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello World' when an emtpy string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
  t.Run("Say hello in spanish", func(t *testing.T) {
    got := Hello("Daniel", "Spanish")
    want := "Hola, Daniel"

    assertCorrectMessage(t, got, want)
  })

  t.Run("Say hello in french", func(t *testing.T) {
    got := Hello("Dior", "French")
    want := "Bonjour, Dior"

    assertCorrectMessage(t, got, want)
  })

  t.Run("Say hello in hindi", func(t *testing.T) {
    got := Hello("Dior", "Hindi")
    want := "Namasthe, Dior"

    assertCorrectMessage(t, got, want)
  })
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
