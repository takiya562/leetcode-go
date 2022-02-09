package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestCountKDifference(t *testing.T) {
	assert.Equal(t, 4, leetcode.CountKDifference([]int{1, 2, 2, 1}, 1))
	assert.Equal(t, 3, leetcode.CountKDifference([]int{3, 2, 1, 5, 4}, 2))
}
