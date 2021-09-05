package substract

import (
	"testing"
	"github.com/NurfitraPujo/go-calculator/test"
)

func TestSubstract(t *testing.T) {

	t.Run("does return substraction of two argument", func(t *testing.T) {
		result := Substract(4, 2)
		expect := 2

		test.AssertEqual(result, expect)
	})

}
