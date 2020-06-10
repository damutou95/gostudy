package main
import "fmt"
func main(){
	nums := []int{2,1,2,2,3,0,4,2}
	val := 2
	fmt.Print(removeElement(nums,val))
}

func removeElement(nums []int, val int) int {
	tag := len(nums)-1
	for i:=len(nums)-1;i>=0;i--{
		if nums[i] == val{
			nums[i],nums[tag]=nums[tag],nums[i]
			tag--
		}
	}
	return tag+1
}
执行用时 :
0 ms
, 在所有 Go 提交中击败了
100.00%
的用户
内存消耗 :
2.1 MB
, 在所有 Go 提交中击败了
9.09%
的用户