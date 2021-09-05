package add

import "testing"

func TestAdd(t *testing.T) {
	t.Run("does return sum of two argument", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
	t.Run("does return sum of array of arguments", func(t *testing.T) {
		sum := Add(2, 2, 6)
		expected := 10

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}
