package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestFindKthNumber(t *testing.T) {
	assert.Equal(t, 10, leetcode.FindKthNumber(13, 2))
	assert.Equal(t, 1, leetcode.FindKthNumber(1, 1))
	assert.Equal(t, 2, leetcode.FindKthNumber(2, 2))
	assert.Equal(t, 2, leetcode.FindKthNumber(10, 3))
	assert.Equal(t, 9, leetcode.FindKthNumber(100, 90))
}
