package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestLeetcode564(t *testing.T) {
	assert.Equal(t, "121", leetcode.NearestPalindromic("123"))
	assert.Equal(t, "0", leetcode.NearestPalindromic("1"))
	assert.Equal(t, "2442", leetcode.NearestPalindromic("2413"))
	assert.Equal(t, "22", leetcode.NearestPalindromic("19"))
	assert.Equal(t, "9", leetcode.NearestPalindromic("10"))
}
