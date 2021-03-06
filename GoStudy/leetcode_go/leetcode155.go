package main
import (
	"fmt"
	"math"
)


type MinStack struct {
	minStack  []int
	stack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{minStack: []int{math.MaxInt64}, stack: []int{}}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	if x <= this.minStack[len(this.minStack)-1]{
		this.minStack = append(this.minStack, x)
	}else{
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	}
}


func (this *MinStack) Pop()  {
	this.minStack = this.minStack[:len(this.minStack)-1]
	this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
