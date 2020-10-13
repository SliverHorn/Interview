package main

import "fmt"

type TreeNode144 struct {
	Val   int
	Left  *TreeNode144
	Right *TreeNode144
}

func main() {
	root := &TreeNode144{
		Val: 1,
		Right: &TreeNode144{
			Val:   2,
			Left:  &TreeNode144{Val: 3},
		},
	}
	fmt.Println(preorderTraversal(root))
}

func preorderTraversal(root *TreeNode144) []int {
	var result []int
	if root != nil {
		result = append(result, root.Val)
		for _, l := range preorderTraversal(root.Left) {
			result = append(result, l)
		}
		for _, r := range preorderTraversal(root.Right) {
			result = append(result, r)
		}
	}
	return result
}
