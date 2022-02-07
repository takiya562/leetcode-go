package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestLongestDiverseString(t *testing.T) {
	assert.Equal(t, "ccaccbcc", leetcode.LongestDiverseString(1, 1, 7))
}
