package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{
		Val:   3,
		Left:  &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	maxDepth(&root)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	fmt.Printf("val:%v,left:%v,right:%v,max:%v, return:%v\n", root.Val ,maxDepth(root.Left), maxDepth(root.Right), max(maxDepth(root.Left), maxDepth(root.Right)), max(maxDepth(root.Left), maxDepth(root.Right)) + 1)
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
