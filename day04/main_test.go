package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input1
var input1 string

func TestOne(t *testing.T) {
	assert.Equal(t, 2, partOne(input1), "they should be equal")
}
