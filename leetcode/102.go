package main

import "fmt"

type TreeNode102 struct {
	Val   int
	Left  *TreeNode102
	Right *TreeNode102
}

func main() {
	var root = &TreeNode102{
		Val: 3,
		Left: &TreeNode102{
			Val:   9,
		},
		Right: &TreeNode102{
			Val:   20,
			Left:  &TreeNode102{
				Val:   15,
			},
			Right:  &TreeNode102{
				Val:   7,
			},
		},
	}
	fmt.Println(levelOrder(root))
}

func levelOrder(root *TreeNode102) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	dfsLevel(root, -1, &result)
	return result
}

func dfsLevel(node *TreeNode102, level int, result *[][]int) {
	if node == nil {
		return
	}

	currLevel := level + 1
	for len(*result) <= currLevel {
		*result = append(*result, []int{})
	}
	(*result)[currLevel] = append((*result)[currLevel], node.Val)
	dfsLevel(node.Left, currLevel, result)
	dfsLevel(node.Right, currLevel, result)
}