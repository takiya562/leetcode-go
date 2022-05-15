package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestMinStickers(t *testing.T) {
	assert.Equal(t, 3, leetcode.MinStickers([]string{"with", "example", "science"}, "thehat"))
	assert.Equal(t, -1, leetcode.MinStickers([]string{"notice", "possible"}, "basicbasic"))
	assert.Equal(t, 1, leetcode.MinStickers([]string{"notice", "possible"}, "notice"))
}
