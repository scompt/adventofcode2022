package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input1
var input1 string

func TestOne(t *testing.T) {
	assert.Equal(t, 3, len(parseInput(input1)), "they should be equal")
	assert.Equal(t, 15, partOne(input1), "they should be equal")
	// assert.Equal(t, 66616, partOne(input2), "they should be equal")
}
