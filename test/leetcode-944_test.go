package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/leetcode"
)

func TestMinDeletionSize(t *testing.T) {
	assert.Equal(t, 1, leetcode.MinDeletionSize([]string{"cba", "daf", "ghi"}))
	assert.Equal(t, 0, leetcode.MinDeletionSize([]string{"a", "b"}))
	assert.Equal(t, 3, leetcode.MinDeletionSize([]string{"zyx", "wvu", "tsr"}))
	assert.Equal(t, 2, leetcode.MinDeletionSize([]string{"rrjk", "furt", "guzm"}))
}
