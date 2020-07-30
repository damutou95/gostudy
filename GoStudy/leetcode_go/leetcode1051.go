package main
import "fmt"
func main(){
	fmt.Print(heightChecker([]int{1,2,3,4,5}))
}
func heightChecker(heights []int) int {
	count := 0
	des := [100]int{}
	for i:=0;i<len(heights);i++{
		des[heights[i]]++
	}
	for j:=0;j<len(heights);j++{
		sumValue1 := 0
		sumValue2 := 0
		for nums:=heights[j];nums>0;nums--{
			sumValue1 += des[nums]
		}
		for nums:=heights[j]-1;nums>0;nums--{
			sumValue2 += des[nums]
		}
		if j < sumValue2 || j >= sumValue1{
			count++
		}
	}
	return count
}
