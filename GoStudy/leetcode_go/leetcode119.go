package main
import "fmt"
func main(){
	fmt.Print(getRow(3))
}
//正向添加
//func getRow(rowIndex int) []int {
//	lastRow := []int{1}
//	currentRow := []int{}
//	if rowIndex == 0{
//		return []int{1}
//	}
//	for i:=1;i<=rowIndex+1;i++{
//		for j:=0;j<i;j++{
//			if j == 0{
//				currentRow = append(currentRow, 1)
//			}else if j == i-1{
//				currentRow = append(currentRow, 1)
//			}else{
//				currentRow = append(currentRow, lastRow[j-1]+lastRow[j])
//			}
//		}
//		lastRow = currentRow
//		currentRow = []int{}
//	}
//	return lastRow
//}
//从尾部添加
func getRow(rowIndex int) []int {
	res := []int{1}
	for i:=1;i<=rowIndex;i++{
		res = append(res, 1)
		for j:=i-1;j>0;j--{
			res[j] = res[j] + res[j-1]
		}
	}
	return res
}
