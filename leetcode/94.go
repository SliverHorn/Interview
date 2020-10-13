package main

import "fmt"

type TreeNode94 struct {
	Val   int
	Left  *TreeNode94
	Right *TreeNode94
}

func main() {
	root := &TreeNode94{
		Val: 1,
		Right: &TreeNode94{
			Val:  2,
			Left: &TreeNode94{Val: 3},
		},
	}
	fmt.Println(inorderTraversal(root))
}

func inorderTraversal(root *TreeNode94) []int {
	var result []int
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode94, result *[]int) {
	if root != nil {
		inorder(root.Left, result)
		*result = append(*result, root.Val)
		inorder(root.Right, result)
	}
}