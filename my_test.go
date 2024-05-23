package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestSum(t *testing.T) {
	total := 5 + 5
	expected := 10
	assert.Equal(t, expected, total, "Sum of 5 and 5 should be 10")
}
