package main
import "fmt"
func main(){
	fmt.Print(movesToMakeZigzag([]int{2,7,10,9,8,9}))
}
func max(a int, b int) int{
	if a > b{
		return a
	}else{
		return b
	}
}
func movesToMakeZigzag(nums []int) int {
	c := 0
	f := 0
	for i:=0;i<len(nums);i++{
		if i % 2 == 0{
			a := 0
			b := 0
			if i > 0 && nums[i] >= nums[i-1]{
				a = nums[i] - nums[i-1] + 1
			}
			if i < len(nums)-1 && nums[i] >= nums[i+1]{
				b = nums[i] - nums[i+1] + 1
			}
			c += max(a, b)
		}else{
			d := 0
			e := 0
			if nums[i] >= nums[i-1]{
				d = nums[i] - nums[i-1] + 1
			}
			if i < len(nums) - 1 && nums[i] >= nums[i+1]{
				e = nums[i] - nums[i+1] + 1
			}
			f += max(d, e)
		}
	}
	return -max(0-c,0-f)
}
