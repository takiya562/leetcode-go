package leetcode

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	dirs := [][]int{{0, 0}, {0, -1}, {0, 1}, {-1, 0}, {-1, -1}, {-1, 1}, {1, 0}, {1, -1}, {1, 1}}
	N := n
	col, row, diag, antiDiag := map[int]int{}, map[int]int{}, map[int]int{}, map[int]int{}
	points := map[int]bool{}
	for _, lamp := range lamps {
		r, c := lamp[0], lamp[1]
		p := r*N + c
		if points[p] {
			continue
		}
		points[p] = true
		col[c]++
		row[r]++
		diag[r-c]++
		antiDiag[r+c]++
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		r, c := query[0], query[1]
		if col[c] > 0 || row[r] > 0 || diag[r-c] > 0 || antiDiag[r+c] > 0 {
			ans[i] = 1
		}
		for _, dir := range dirs {
			nr, nc := r+dir[0], c+dir[1]
			if nr < 0 || nc < 0 || nr >= n || nc >= n {
				continue
			}
			p := nr*N + nc
			if points[p] {
				points[p] = false
				col[nc]--
				row[nr]--
				diag[nr-nc]--
				antiDiag[nr+nc]--
			}
		}
	}
	return ans
}

func GridIllumination(n int, lamps [][]int, queries [][]int) []int {
	return gridIllumination(n, lamps, queries)
}
