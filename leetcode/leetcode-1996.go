package leetcode

import "sort"

func NumberOfWeakCharacters(properties [][]int) int {
	return numberOfWeakCharacters(properties)
}

func numberOfWeakCharacters(properties [][]int) (ans int) {
	sort.Slice(properties, func(i, j int) bool {
		p, q := properties[i], properties[j]
		return p[0] < q[0] || p[0] == q[0] && p[1] > q[1]
	})
	st := []int{}
	for _, p := range properties {
		for len(st) > 0 && st[len(st)-1] < p[1] {
			st = st[:len(st)-1]
			ans++
		}
		st = append(st, p[1])
	}
	return
}
