package main
import "fmt"
func main(){
	fmt.Print(generate(5))
}

func generate(numRows int) [][]int {
	if numRows==0{
		return [][]int{}
	}
	triangle := [][]int{[]int{1}}
	lastNums := 0
	nowNums := 0
	for i:=0;i<numRows-1;i++{
		triangle = append(triangle, []int{})
		for nowNums<lastNums+1{
			if nowNums==0{
				triangle[i+1] = append(triangle[i+1],1)
			}else {
				triangle[i+1] = append(triangle[i+1],triangle[i][nowNums-1]+triangle[i][nowNums])
			}
			nowNums++
		}
		triangle[i+1] = append(triangle[i+1],1)
		lastNums++
		nowNums = 0
	}
	return triangle
}
