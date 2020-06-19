package main

import (
	"fmt"
)
func main(){
	x := 1
	fmt.Print(mySqrt(x))
}
func mySqrt(x int) int {
	left := 0
	right := x
	for left<right{
		mid := left + (right-left)/2
		if mid*mid>x{
			right = mid
		}else if (mid+1)*(mid+1)>x{
			return mid
		}else if (mid+1)*(mid+1)==x{
			return mid+1
		}else{
			left = mid
		}
	}
	return 0
}
