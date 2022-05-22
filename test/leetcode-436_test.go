package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestFindRightInterval(t *testing.T) {
	assert.Equal(t, []int{-1}, leetcode.FindRightInterval([][]int{{1, 2}}))
	assert.Equal(t, []int{-1, 0, 1}, leetcode.FindRightInterval([][]int{{3, 4}, {2, 3}, {1, 2}}))
	assert.Equal(t, []int{-1, 2, -1}, leetcode.FindRightInterval([][]int{{1, 4}, {2, 3}, {3, 4}}))
	assert.Equal(t, []int{3, 2, 3, -1}, leetcode.FindRightInterval([][]int{{1, 4}, {2, 3}, {3, 4}, {4, 5}}))
	assert.Equal(t, []int{1, 3, 0, -1}, leetcode.FindRightInterval([][]int{{1, 2}, {2, 3}, {0, 1}, {3, 4}}))
	assert.Equal(t, []int{2, 3, -1, -1}, leetcode.FindRightInterval([][]int{{-100, -98}, {-99, -97}, {-98, -96}, {-97, -95}}))
}
