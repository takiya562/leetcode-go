package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestCanIWin(t *testing.T) {
	assert.Equal(t, false, leetcode.CanIWin(10, 11))
	assert.Equal(t, true, leetcode.CanIWin(10, 0))
	assert.Equal(t, true, leetcode.CanIWin(10, 1))
}
