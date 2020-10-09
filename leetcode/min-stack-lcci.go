package main

import "fmt"

type MinStack struct {
	l int
	List, Min []int

}


/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
	this.List = append(this.List, x)
	if this.l == 0 {
		this.Min = append(this.Min, x)
	} else {
		min := this.GetMin()
		if x < min {
			this.Min = append(this.Min, x)
		} else {
			this.Min = append(this.Min, min)
		}
	}
	this.l++
}


func (this *MinStack) Pop() {
	this.l--
	this.Min = this.Min[:this.l]
	this.List = this.List[:this.Length()-1]
}


func (this *MinStack) Top() int {
	if this.Length() == 0 {
		return 0
	}
	return this.List[this.Length()-1]
}


func (this *MinStack) GetMin() int {
	if this.Length() == 0 {
		return 0
	}
	return this.Min[this.l-1]
}

func (this *MinStack) Length() int {
	return len(this.List)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	x := 0
	obj := MinStackConstructor()
	obj.Push(x)
	obj.Pop()
	a := obj.Top()
	b := obj.GetMin()
	fmt.Println(a, b)
}