package main

type MyStack struct {
	List []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.List = append(this.List, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Length() == 0 {
		return 0
	}
	last := this.List[this.Length()-1]
	this.List = this.List[:this.Length()-1]
	return last
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.List[this.Length()-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Length() == 0
}

/** Returns whether the stack is empty. */
func (this *MyStack) Length() int {
	return len(this.List)
}



/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
