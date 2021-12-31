package main

import (
	"fmt"
	"math"
)

func isNStraightHand(hand []int, groupSize int) bool {
	l := len(hand)
	if l%groupSize != 0 {
		return false
	}
	max, min := hand[0], hand[0]
	m := make(map[int]int)
	for _, v := range hand {
		min = int(math.Min(float64(min), float64(v)))
		max = int(math.Max(float64(max), float64(v)))
		m[v]++
	}
	for {
		for m[min] == 0 {
			if min > max {
				return true
			}
			min++
		}
		for i := min; i < min+groupSize; i++ {
			if m[i] == 0 {
				return false
			}
			m[i]--
		}
	}
}

func main() {
	hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	fmt.Println(isNStraightHand(hand, 3))
}
