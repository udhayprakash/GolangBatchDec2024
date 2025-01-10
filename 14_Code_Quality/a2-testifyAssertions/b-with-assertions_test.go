package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go get "github.com/stretchr/testify/assert"
func Add(a, b int) int {
	return a + b
}

func TestAddAssertions(t *testing.T) {
	result := Add(2, 3)
	assert.Equal(t, 5, result, "2 + 3 should equal 5")

	result = Add(-1, 1)
	assert.Equal(t, 10, result, "-1 + 1 should equal 0")

	result = Add(0, 0)
	assert.Equal(t, 0, result, "0 + 0 should equal 0")
}

func TestNil(t *testing.T) {
	var ptr *int
	assert.Nil(t, ptr, "ptr should be nil")

	ptr = new(int)
	assert.NotNil(t, ptr, "ptr should not be nil")
}
