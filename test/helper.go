package test

import "testing"

func PrintError(got, want int, t *testing.T) {
	t.Errorf("expected '%d' but got '%d'", want, got)
}

func AssertEqual(got, want int) {
	if got != want {
		PrintError(got, want, &testing.T{})
	}
}
