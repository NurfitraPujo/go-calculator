package multiply

import (
	"testing"

	"github.com/NurfitraPujo/go-calculator/test"
)

func TestMultiply(t *testing.T) {
	t.Run("does return multiplication result of two arguments", func(t *testing.T) {
		result := Multiply(3, 4)
		expect := 12

		test.AssertEqual(t, result, expect)
	})
}
