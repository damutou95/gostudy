package main


//直接遍历复杂度O(n)
//func missingNumber(nums []int) int {
//	for i:=0;i<len(nums);i++{
//		if nums[i] != i{
//			return i
//		}
//	}
//	return len(nums)
//}

import "fmt"
func main(){
	fmt.Print(missingNumber([]int{0,1,3}))
}

func missingNumber(nums []int) int {
	i := 0
	j := len(nums) - 1
	for i <= j{
		m := i + (j -i) / 2
		if nums[m] == m{
			i = m + 1
		}else{
			j = m - 1
		}
	}
	return i
}
