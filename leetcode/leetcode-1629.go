package leetcode

func slowestKey(releaseTimes []int, keysPressed string) byte {
	max, pre := 0, 0
	res := byte(keysPressed[0])
	var t int
	for i, release := range releaseTimes {
		t = release - pre
		if t > max || t == max && res < keysPressed[i] {
			res = keysPressed[i]
			max = t
		}
		pre = release
	}
	return res
}
