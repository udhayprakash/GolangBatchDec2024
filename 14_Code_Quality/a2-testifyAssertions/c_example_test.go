package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	result := 2 + 2
	assert.Equal(t, 4, result, "2 + 2 should equal 4")
}

func TestNotEqual(t *testing.T) {
	result := 2 + 3
	assert.NotEqual(t, 4, result, "2 + 3 should not equal 4")
}

func TestNil2(t *testing.T) {
	var ptr *int
	assert.Nil(t, ptr, "ptr should be nil")
}

func TestNotNil(t *testing.T) {
	ptr := new(int)
	assert.NotNil(t, ptr, "ptr should not be nil")
}

func TestBoolean(t *testing.T) {
	isTrue := true
	assert.True(t, isTrue, "isTrue should be true")

	isFalse := false
	assert.False(t, isFalse, "isFalse should be false")
}

func someFunctionThatReturnsError() err {
	return error.New("ass")
}
func TestError(t *testing.T) {
	err := someFunctionThatReturnsError()
	assert.Error(t, err, "someFunctionThatReturnsError should return an error")
}

func TestNoError(t *testing.T) {
	err := someFunctionThatDoesNotReturnError()
	assert.NoError(t, err, "someFunctionThatDoesNotReturnError should not return an error")
}

func TestPanic(t *testing.T) {
	assert.Panics(t, func() {
		panic("this should panic")
	}, "The function should panic")
}

func TestSliceEquality(t *testing.T) {
	expected := []int{1, 2, 3}
	actual := []int{1, 2, 3}
	assert.Equal(t, expected, actual, "Slices should be equal")
}
// assignment : compare two maps, write an assertion for it

