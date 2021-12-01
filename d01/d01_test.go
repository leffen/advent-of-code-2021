package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD01(t *testing.T) {
	nums := importFileAsIntA("test_data.txt")
	assert.Equal(t, 10, len(nums))
	cnt := count(nums)
	assert.Equal(t, int64(7), cnt)
}
