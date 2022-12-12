package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Enzo")
	want := "Hello, Enzo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
