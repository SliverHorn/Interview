package main

import (
	"fmt"
	"sort"
)

func main() {
	var A = []int{-4,-1,0,3,10}
	fmt.Println(sortedSquares(A))

}

func sortedSquares(A []int) []int {
	l := len(A)
	if l == 0 {
		return []int{}
	}
	result := make([]int, l)
	for i, v := range A {
		result[i] = v*v
	}
	sort.Ints(result)
	return result
}