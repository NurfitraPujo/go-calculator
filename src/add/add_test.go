package add

import (
	"testing"

	"github.com/NurfitraPujo/go-calculator/test"
)

func TestAdd(t *testing.T) {
	t.Run("does return sum of two argument", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		test.AssertEqual(t, sum, expected)
	})
	t.Run("does return sum of array of arguments", func(t *testing.T) {
		sum := Add(2, 2, 6)
		expected := 10

		test.AssertEqual(t, sum, expected)
	})
}
