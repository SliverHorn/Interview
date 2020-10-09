package main

import "fmt"

func main() {
	nums := []int{2,0,2,1,1,0}
	sortColors(nums)
}

func sortColors(nums []int)  {
	if len(nums) == 0 {
		return
	}
	a,b,c := 0,0,0
	for _, num := range nums {
		if num == 0 {
			nums[c] = 2
			c++
			nums[b] = 1
			b++
			nums[a] = 0
			a++
		} else if num == 1 {
			nums[c] = 2
			c++
			nums[b] = 1
			b++
		} else if num == 2 {
			c++
		}
	}
	fmt.Println(nums)

}
