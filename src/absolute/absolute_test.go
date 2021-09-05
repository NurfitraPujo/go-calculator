package absolute

import "testing"

func TestAbsolute(t *testing.T) {

	t.Run("does return absolute value of given number", func(t *testing.T) {
		result := Absolute(-2)
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
