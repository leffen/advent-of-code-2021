package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func xTestD05(t *testing.T) {
	tx := loadData("test_data.txt", false, false)
	assert.NotNil(t, tx)
	num := tx.NumOverlapPoints()
	assert.Equal(t, 5, num)
}

func TestD05p2(t *testing.T) {
	tx := loadData("test_data.txt", true, true)
	assert.NotNil(t, tx)
	num := tx.NumOverlapPoints()
	fmt.Println("----- RESULT ------")
	tx.Show()

	numX := 0
	for _, c := range tx.cells {
		//	fmt.Printf("  %s (%d,%d) %d\n", c.ID(), c.x, c.y, c.overlaps)
		if c.overlaps > 0 {
			numX++
		}
	}

	fmt.Printf("Num points %d\n", numX)

	assert.Equal(t, 12, num)
}
