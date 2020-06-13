package main
import "fmt"
func main(){
	fmt.Print(searchInsert([]int{1,4,6,8,9},4))
}
func searchInsert(nums []int, target int) int {
	if len(nums) == 0{
		return 0
	}
	left := 0
	right := len(nums)-1

	for left<=right{
		mid := left + (right-left)/2
		if target == nums[mid]{
			return mid
		}else if target<nums[mid]{
			right = mid - 1
		}else{
			left = mid + 1
		}
	}
	return left

}