package leetcode

import "math"

type DetectSquares struct {
	graph [][][2]int
}

func constructor() DetectSquares {
	return DetectSquares{
		graph: make([][][2]int, 1001),
	}
}

func (this *DetectSquares) Add(point []int) {
	x := point[0]
	y := point[1]
	this.graph[x] = append(this.graph[x], [2]int{x, y})
}

func (this *DetectSquares) Count(point []int) int {
	x := point[0]
	y := point[1]
	count := 0
	for _, p := range this.graph[x] {
		if p[1] == y {
			continue
		}
		l := int(math.Abs(float64(p[1] - y)))
		c1, c2 := this.find(p[0]-l, p[1], y)
		c3, c4 := this.find(p[0]+l, p[1], y)
		count += c1*c2 + c3*c4
	}
	return count
}

func (this *DetectSquares) find(x, y1, y2 int) (int, int) {
	if x < 0 || x > 1000 {
		return 0, 0
	}
	c1, c2 := 0, 0
	for _, p := range this.graph[x] {
		if p[1] == y1 {
			c1++
		}
		if p[1] == y2 {
			c2++
		}
	}
	return c1, c2
}
