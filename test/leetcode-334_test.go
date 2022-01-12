package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestLeetcode334(t *testing.T) {
	assert.True(t, leetcode.IncreasingTriplet([]int{1, 2, 3, 4, 5}))
	assert.True(t, leetcode.IncreasingTriplet([]int{2, 1, 5, 0, 4, 6}))
	assert.True(t, leetcode.IncreasingTriplet([]int{1, 5, 0, 4, 1, 3}))
	assert.True(t, leetcode.IncreasingTriplet([]int{4, 5, 2147483647, 1, 2}))
	assert.False(t, leetcode.IncreasingTriplet([]int{5, 4, 3, 2, 1}))
}
