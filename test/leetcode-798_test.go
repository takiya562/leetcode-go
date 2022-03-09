package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestBestRotation(t *testing.T) {
	assert.Equal(t, 3, leetcode.BestRotation([]int{2, 3, 1, 4, 0}))
	assert.Equal(t, 0, leetcode.BestRotation([]int{1, 3, 0, 2, 4}))
}
