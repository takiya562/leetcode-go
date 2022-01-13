package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestDominantIndex(t *testing.T) {
	assert.Equal(t, 1, leetcode.DominantIndex([]int{3, 6, 1, 0}))
	assert.Equal(t, -1, leetcode.DominantIndex([]int{1, 2, 3, 4}))
	assert.Equal(t, 0, leetcode.DominantIndex([]int{1}))
	assert.Equal(t, -1, leetcode.DominantIndex([]int{0, 0, 3, 2}))
}
