package leetcode

import (
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor449() Codec {
	return Codec{}
}

func preOrder(root *TreeNode) string {
	if root == nil {
		return ""
	}
	arr := []string{}
	s := []*TreeNode{}
	cur := root
	for len(s) != 0 || cur != nil {
		for cur != nil {
			arr = append(arr, strconv.Itoa(cur.Val))
			s = append(s, cur)
			cur = cur.Left
		}
		cur = s[len(s)-1].Right
		s = s[:len(s)-1]
	}
	return strings.Join(arr, " ")
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return preOrder(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, " ")
	var construct func(int, int) *TreeNode
	construct = func(lower, upper int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		val, _ := strconv.Atoi(arr[0])
		if val < lower || val > upper {
			return nil
		}
		arr = arr[1:]
		return &TreeNode{
			Val:   val,
			Left:  construct(lower, val),
			Right: construct(val, upper),
		}
	}
	return construct(math.MinInt32, math.MaxInt32)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
