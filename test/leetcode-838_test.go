package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestPushDominoes(t *testing.T) {
	assert.Equal(t, "RR.L", leetcode.PushDominoes("RR.L"))
	assert.Equal(t, "LL.RR.LLRRLL..", leetcode.PushDominoes(".L.R...LR..L.."))
}
