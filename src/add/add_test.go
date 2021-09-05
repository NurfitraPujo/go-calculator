package add

import "testing"

func TestAdd(t *testing.T) {
	t.Run("does return sum of two argument", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertEqual(sum, expected)
	})
	t.Run("does return sum of array of arguments", func(t *testing.T) {
		sum := Add(2, 2, 6)
		expected := 10

		assertEqual(sum, expected)
	})
}

func printError(got, want int, t *testing.T) {
	t.Errorf("expected '%d' but got '%d'", want, got)
}

func assertEqual(got, want int) {
	if got != want {
		printError(got, want, &testing.T{})
	}
}
