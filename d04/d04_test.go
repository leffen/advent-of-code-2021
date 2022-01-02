package main

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestD4(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	g := loadData("test_data.txt")
	assert.NotNil(t, g)
	assert.Equal(t, 27, len(g.Draws))
	assert.Equal(t, 3, len(g.Boards))
	for _, b := range g.Boards {
		assert.Equal(t, 25, len(b.cells))
	}

	score := g.DoDraw()
	assert.Equal(t, 4512, score)
	s2 := g.DoDraw2()
	assert.Equal(t, 1924, s2)

	//assert.True(t, false)
}
