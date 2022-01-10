package test

import (
	"testing"

	"github.com/takiya562/leetcode-go/leetcode"
)

func TestLeetCode306(t *testing.T) {
	if res := leetcode.IsAdditiveNumber("112358"); !res {
		t.Errorf("IsAdditiveNumber(%s)=%t want true", "112358", res)
	}
	if res := leetcode.IsAdditiveNumber("1023"); res {
		t.Errorf("IsAdditiveNumber(%s)=%t want false", "1023", res)
	}
	if res := leetcode.IsAdditiveNumber("101"); !res {
		t.Errorf("IsAdditiveNumber(%s)=%t want true", "101", res)
	}
	if res := leetcode.IsAdditiveNumber("211738"); !res {
		t.Errorf("IsAdditiveNumber(%s)=%t want true", "211738", res)
	}
	if res := leetcode.IsAdditiveNumber("0235813"); res {
		t.Errorf("IsAdditiveNumber(%s)=%t want false", "0235813", res)
	}
	if res := leetcode.IsAdditiveNumber("000"); !res {
		t.Errorf("IsAdditiveNumber(%s)=%t want true", "000", res)
	}
	if res := leetcode.IsAdditiveNumber("12122436"); !res {
		t.Errorf("IsAdditiveNumber(%s)=%t want true", "12122436", res)
	}
	if res := leetcode.IsAdditiveNumber("198019823962"); !res {
		t.Errorf("IsAdditiveNumber(%s)=%t want true", "198019823962", res)
	}
}
