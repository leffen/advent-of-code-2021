package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	lines, err := importFile("test_data.txt")
	assert.Nil(t, err)
	assert.Equal(t, 12, len(lines))

	gamma := calcGamma(lines)
	assert.Equal(t, int64(22), gamma)

	epsi := calcEpsilon(lines)
	assert.Equal(t, int64(9), epsi)

}

func TestPart2(t *testing.T) {
	lines, err := importFile("test_data.txt")
	assert.Nil(t, err)
	assert.Equal(t, 12, len(lines))

	oxygen := calcOxygen(lines)
	assert.Equal(t, int64(23), oxygen)

	scrubber := calcScrubber(lines)
	assert.Equal(t, int64(10), scrubber)

}
