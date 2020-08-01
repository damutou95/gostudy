package main
import "fmt"
func main(){
	fmt.Print(uniquePaths(7,3))
}
//递归，超时
//func uniquePaths(m int, n int) int {
//	if m==1{
//		return 1
//	}
//	if n==1{
//		return 1
//	}
//	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
//}

func uniquePaths(m int, n int) int {
	temp := []int{}
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if i==0 || j==0{
				temp = append(temp, 1)
			} else{
				temp = append(temp, temp[(i-1)*n+j]+temp[i*n+j-1])
			}
		}
	}
	return temp[len(temp)-1]
}