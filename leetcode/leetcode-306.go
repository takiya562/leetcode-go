package leetcode

import "strconv"

func IsAdditiveNumber(num string) bool {
	var n1, n2 int
	for i := 2; i <= 2*len(num)/3+1; i++ {
		for j := 1; j < i; j++ {
			if i-j > 1 && num[j] == '0' || j > 1 && num[0] == '0' {
				continue
			}
			n1, _ = strconv.Atoi(num[0:j])
			n2, _ = strconv.Atoi(num[j:i])
			l := len(strconv.Itoa(n1 + n2))
			if i+l > len(num) {
				continue
			} else if n3, _ := strconv.Atoi(num[i : i+l]); n3 == n1+n2 && isAdditiveNumber0(num[i+l:], n2, n3) {
				return true
			}
		}
	}
	return false
}

func isAdditiveNumber0(num string, n1, n2 int) bool {
	if len(num) == 0 {
		return true
	}
	l := len(strconv.Itoa(n1 + n2))
	if l > len(num) || l > 1 && num[0] == '0' {
		return false
	} else if n3, _ := strconv.Atoi(num[0:l]); n3 != (n1 + n2) {
		return false
	} else {
		return isAdditiveNumber0(num[l:], n2, n3)
	}
}
