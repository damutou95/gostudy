package main

import (
	"fmt"
)
func main(){
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Print(maxSubArray(nums))
}
func maxSubArray(nums []int) int {
	temp := nums[0]
	res := temp
	for _,v := range nums[1:] {
		if temp + v > v{
			temp = temp + v
		}else{
			temp = v
		}
		if res<temp{
			res = temp
		}
	}
return res

}