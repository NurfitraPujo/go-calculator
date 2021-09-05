package divide

import (
	"testing"

	"github.com/NurfitraPujo/go-calculator/test"
)

func TestDivide(t *testing.T) {
	t.Run("does return division result of two arguments", func(t *testing.T) {
		result := Divide(4, 4)
		expect := 1

		test.AssertEqual(t, result, expect)
	})
}
