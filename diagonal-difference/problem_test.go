package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	input := [][]int32{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}
	res := diagonalDifference(input)
	assert.Equal(t, res, int32(0))
}

func TestProblem2(t *testing.T) {
	input := [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}
	res := diagonalDifference(input)
	assert.Equal(t, res, int32(15))
}
