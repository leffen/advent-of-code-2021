package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	cmds, err := loadDirectionsFile("test_data.txt")
	assert.Nil(t, err)
	assert.Equal(t, 6, len(cmds))
	num := walk(cmds)
	assert.Equal(t, int64(150), num)
}

func TestLoadVer2(t *testing.T) {
	cmds, err := loadDirectionsFile("test_data.txt")
	assert.Nil(t, err)
	assert.Equal(t, 6, len(cmds))
	num := walk2(cmds)
	assert.Equal(t, int64(900), num)
}
