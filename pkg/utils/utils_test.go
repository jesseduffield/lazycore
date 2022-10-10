package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMin is a function.
func TestMin(t *testing.T) {
	type scenario struct {
		a        int
		b        int
		expected int
	}

	scenarios := []scenario{
		{
			1,
			1,
			1,
		},
		{
			1,
			2,
			1,
		},
		{
			2,
			1,
			1,
		},
	}

	for _, s := range scenarios {
		assert.EqualValues(t, s.expected, Min(s.a, s.b))
	}
}

func TestGetLazyRootDirectory(t *testing.T) {
	assert.NotPanics(t, func() {
		GetLazyRootDirectory()
	})
}
