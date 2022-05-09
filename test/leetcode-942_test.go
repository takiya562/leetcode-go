package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestDiStringMatch(t *testing.T) {
	assert.Equal(t, []int{0, 4, 1, 3, 2}, leetcode.DiStringMatch("IDID"))
	assert.Equal(t, []int{0, 1, 2, 3}, leetcode.DiStringMatch("III"))
	assert.Equal(t, []int{3, 2, 0, 1}, leetcode.DiStringMatch("DDI"))
}
