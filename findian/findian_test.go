package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasIan_GOOD(t *testing.T) {
	good := []string{"ian", "Ian", "iuiygaygn", "I d skd a efju N"}

	for _, g := range good {
		assert.True(t, HasIan(g))
	}

	bad := []string{"ihhhhhn", "ina", "xian"}

	for _, b := range bad {
		assert.False(t, HasIan(b))
	}
}
