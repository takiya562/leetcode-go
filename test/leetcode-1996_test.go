package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestNumberOfWeakCharacters(t *testing.T) {
	assert.Equal(t, 0, leetcode.NumberOfWeakCharacters([][]int{{5, 5}, {6, 3}, {3, 6}}))
	assert.Equal(t, 1, leetcode.NumberOfWeakCharacters([][]int{{2, 2}, {3, 3}}))
	assert.Equal(t, 1, leetcode.NumberOfWeakCharacters([][]int{{1, 5}, {10, 4}, {4, 3}}))
}
