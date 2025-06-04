package main

import "testing"

func TestHelloWorld(t *testing.T) {
    got := HelloWorld("Fermatix")
    want := "Hello, Fermatix!"

    if got != want {
        t.Errorf("HelloWorld(\"Fermatix\") = %q; want %q", got, want)
    }
}

func TestHelloWorld_EmptyName(t *testing.T) {
    got := HelloWorld("")
    want := "Hello, !"

    if got != want {
        t.Errorf("HelloWorld(\"\") = %q; want %q", got, want)
    }
}
