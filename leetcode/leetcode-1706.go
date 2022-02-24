package leetcode

const RIGHT = 1
const LEFT = -1

func findBall(grid [][]int) []int {
	r := len(grid)
	c := len(grid[0])
	ans := []int{}
	for i := 0; i < c; i++ {
		ball := i
		for j := 0; j < r; j++ {
			cur := grid[j][ball]
			if cur == RIGHT {
				if ball >= c-1 || grid[j][ball+1] == LEFT {
					ball = -1
					break
				}
				ball++
			}
			if cur == LEFT {
				if ball <= 0 || grid[j][ball-1] == RIGHT {
					ball = -1
					break
				}
				ball--
			}
		}
		ans = append(ans, ball)
	}
	return ans
}

func FindBall(grid [][]int) []int {
	return findBall(grid)
}
