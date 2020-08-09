package main
import "fmt"

func main(){
	//a := [][]int{{0,0,1,0,0,0,0,1,0,0,0,0,0},
	//{0,0,0,0,0,0,0,1,1,1,0,0,0},
	//{0,1,1,0,1,0,0,0,0,0,0,0,0},
	//{0,1,0,0,1,1,0,0,1,0,1,0,0},
	//{0,1,0,0,1,1,0,0,1,1,1,0,0},
	//{0,0,0,0,0,0,0,0,0,0,1,0,0},
	//{0,0,0,0,0,0,0,1,1,1,0,0,0},
	//{0,0,0,0,0,0,0,1,1,0,0,0,0}}
	a := [][]int{
		{0,1},
		{1,1},
	}
	fmt.Print(maxAreaOfIsland(a))
}

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0{
		return 0
	}
	if len(grid[0]) == 0{
		return 0
	}
	temp := 0
	res := 0


	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++ {
			temp = helper(grid, i, j)
			if temp > res{
				res = temp
			}
		}
	}
	return res
}


func helper(temp [][]int, m int, n int) int{
	if len(temp) == 0{
		return 0
	}
	if len(temp[0]) == 0{
		return 0
	}
	res := 0
	if m>=0 && m<len(temp) && n>=0 && n<len(temp[0]){
		if temp[m][n]==1{
			temp[m][n] = 0
			res += 1
			if m>=1{
				res += helper(temp, m-1, n)
			}
			if m<len(temp)-1{
				res += helper(temp, m+1, n)
			}
			if n>=1{
				res += helper(temp, m, n-1)
			}
			if n<len(temp[0])-1{
				res += helper(temp, m, n+1)
			}
		}

	}
	return res

}