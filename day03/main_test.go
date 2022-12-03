package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input1
var input1 string

func TestOne(t *testing.T) {
	assert.Equal(t, 1, scoreChar('a'))
	assert.Equal(t, 26, scoreChar('z'))
	assert.Equal(t, 27, scoreChar('A'))
	assert.Equal(t, 52, scoreChar('Z'))

	assert.Equal(t, 6, len(parseInput(input1)), "they should be equal")
	assert.Equal(t, 157, partOne(input1), "they should be equal")
	// assert.Equal(t, 66616, partOne(input2), "they should be equal")
}

func TestFindDupe(t *testing.T) {
	r := rucksack{"asdf", "fghj", 4}
	dupe := r.findDupe()
	assert.Equal(t, 'f', dupe)
}
