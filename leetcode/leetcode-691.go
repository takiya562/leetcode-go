package leetcode

func minStickers(stickers []string, target string) int {
	n := len(stickers)
	end := 1<<len(target) - 1
	sCnt := make([][]int, n)
	c2S := make(map[rune]bool)
	for i, s := range stickers {
		cnt := make([]int, 26)
		for _, c := range s {
			cnt[c-'a']++
			if !c2S[c] {
				c2S[c] = true
			}
		}
		sCnt[i] = cnt
	}
	for _, c := range target {
		if !c2S[c] {
			return -1
		}
	}
	searched := make(map[int]bool)
	states := []int{0}
	nextState := func(state, index int) int {
		cnt := sCnt[index]
		for i := range target {
			if state&(1<<i) != 0 || cnt[target[i]-'a'] == 0 {
				continue
			}
			cnt[target[i]-'a']--
			state |= 1 << i
		}
		return state
	}
	finish := func(state int) bool {
		return state == end
	}
	step := 0
	for len(states) > 0 {
		next := []int{}
		for _, state := range states {
			if finish(state) {
				return step
			}
			for i := range stickers {
				ns := nextState(state, i)
				if isSearch, ok := searched[ns]; !ok || !isSearch {
					searched[ns] = true
					next = append(next, ns)
				}
			}
		}
		step++
		states = next
	}
	return step
}

func MinStickers(stickers []string, target string) int {
	return minStickers(stickers, target)
}
