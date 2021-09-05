package substract

import "testing"

func TestSubstract(t *testing.T) {

	t.Run("does return substraction of two argument", func(t *testing.T) {
		result := Substract(4, 2)
		expect := 2

		assertEqual(result, expect)
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
