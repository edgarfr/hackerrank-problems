package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	input := []int32{3, 2, 1, 3}
	res := birthdayCakeCandles(input)
	assert.Equal(t, res, int32(2))
}

func TestProblem2(t *testing.T) {
	input := []int32{4, 4, 1, 3}
	res := birthdayCakeCandles(input)
	assert.Equal(t, res, int32(2))
}
