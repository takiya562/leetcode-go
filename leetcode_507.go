package main

import "math"

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	n := 1
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			n += i
			n += num / i
		}
	}
	return num == n
}

func main() {

}
