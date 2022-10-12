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

func TestClamp(t *testing.T) {
	tests := []struct {
		name string
		x    int
		min  int
		max  int
		want int
	}{
		{
			"successX",
			5,
			1,
			10,
			5,
		},
		{
			"successMin",
			-5,
			1,
			10,
			1,
		},
		{
			"successMax",
			15,
			1,
			10,
			10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Clamp(tt.x, tt.min, tt.max))
		})
	}
}

func TestGetLazyRootDirectory(t *testing.T) {
	assert.NotPanics(t, func() {
		GetLazyRootDirectory()
	})
}
