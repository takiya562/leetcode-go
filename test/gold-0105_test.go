package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takiya562/leetcode-go/gold"
)

func TestOneEditAway(t *testing.T) {
	assert.Equal(t, true, gold.OneEditAway("pale", "ple"))
	assert.Equal(t, false, gold.OneEditAway("pales", "pal"))
	assert.Equal(t, true, gold.OneEditAway("pales", "palee"))
	assert.Equal(t, true, gold.OneEditAway("ple", "pale"))
}
