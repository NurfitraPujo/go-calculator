package absolute

import (
	"testing"

	"github.com/NurfitraPujo/go-calculator/test"
)

func TestAbsolute(t *testing.T) {

	t.Run("does return absolute value of given number", func(t *testing.T) {
		result := Absolute(-2)
		expect := 2

		test.AssertEqual(result, expect)
	})

}
