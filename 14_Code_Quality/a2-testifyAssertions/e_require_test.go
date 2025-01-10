package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestRequire(t *testing.T) {
	result := 2 + 2
	require.Equal(t, 4, result, "2 + 2 should equal 4") // Test stops here if this fails
	t.Log("This will only run if the above assertion passes")
}

// require package works like assert, but it immediately fails the test
//  (calls t.FailNow()) if the assertion fails

