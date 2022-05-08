package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestFindDuplicates(t *testing.T) {
	assert.Equal(t, []int{2, 3}, leetcode.FindDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
