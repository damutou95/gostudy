package main
import "fmt"
func main(){
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Print(removeDuplicates(nums))
}
//func removeDuplicates(nums []int) int {
//	tag := 0
//	cur := 0
//	for cur < len(nums){
//		if nums[tag] != nums[cur]{
//			nums[tag+1] = nums[cur]
//			tag++
//		}
//		cur++
//	}
//	return tag+1
//}
func removeDuplicates(nums []int) int {
	tag := 0
	for _,v := range nums{
		if nums[tag] != v{
			nums[tag+1] = v
			tag++
		}
	}
	return tag+1
}