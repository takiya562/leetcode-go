package leetcode

import "container/heap"

type pair struct {
	letter byte
	count  int
}

type pairHeap []*pair

func (h pairHeap) Len() int {
	return len(h)
}

func (h pairHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h pairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *pairHeap) Push(x interface{}) {
	*h = append(*h, x.(*pair))
}

func (h *pairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func longestDiverseString(a int, b int, c int) string {
	h := &pairHeap{}
	if a > 0 {
		*h = append(*h, &pair{letter: 'a', count: a})
	}
	if b > 0 {
		*h = append(*h, &pair{letter: 'b', count: b})
	}
	if c > 0 {
		*h = append(*h, &pair{letter: 'c', count: c})
	}
	heap.Init(h)
	ans := []byte{}
	for len(*h) > 0 {
		cur := heap.Pop(h).(*pair)
		n := len(ans)
		if n >= 2 && ans[n-1] == cur.letter && ans[n-2] == cur.letter {
			if len(*h) == 0 {
				break
			}
			next := heap.Pop(h).(*pair)
			ans = append(ans, next.letter)
			next.count--
			if next.count != 0 {
				heap.Push(h, next)
			}
			heap.Push(h, cur)
		} else {
			ans = append(ans, cur.letter)
			cur.count--
			if cur.count != 0 {
				heap.Push(h, cur)
			}
		}
	}
	return string(ans)
}

func LongestDiverseString(a, b, c int) string {
	return longestDiverseString(a, b, c)
}
