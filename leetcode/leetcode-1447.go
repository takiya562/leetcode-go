package leetcode

import "strconv"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func simplifiedFractions(n int) []string {
	ans := []string{}
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if gcd(i, j) == 1 {
				ans = append(ans, strconv.Itoa(i)+"/"+strconv.Itoa(j))
			}
		}
	}
	return ans
}

func SimplifiedFractions(n int) []string {
	return simplifiedFractions(n)
}
