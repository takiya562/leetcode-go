package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestLeetCode2047(t *testing.T) {
	assert.Equal(t, 4, leetcode.CountValidWords(". ! 7hk  al6 l! aon49esj35la k3 7u2tkh  7i9y5  !jyylhppd et v- h!ogsouv 5"))
	assert.Equal(t, 7, leetcode.CountValidWords(" , wn xhx5e2  9k !  fs uc5jc  u3  f5 lj x mkkouek, .g"))
	assert.Equal(t, 3, leetcode.CountValidWords("cat and  dog"))
	assert.Equal(t, 0, leetcode.CountValidWords("!this  1-s b8d!"))
	assert.Equal(t, 5, leetcode.CountValidWords("alice and  bob are playing stone-game10"))
	assert.Equal(t, 6, leetcode.CountValidWords("he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."))
	assert.Equal(t, 1, leetcode.CountValidWords("!"))
	assert.Equal(t, 1, leetcode.CountValidWords("a"))
	assert.Equal(t, 1, leetcode.CountValidWords(" o6 t"))
}
