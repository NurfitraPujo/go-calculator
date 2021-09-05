package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func notEqualMessage(got, want int) string {
	return fmt.Sprintf("expected '%d' but got '%d'", want, got)
}

func AssertEqual(t *testing.T, got, want int) {
	assert.Equal(t, got, want, notEqualMessage(got, want))
}
