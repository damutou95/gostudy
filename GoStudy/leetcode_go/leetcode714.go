package main
import (
	"fmt"
	"math"
)
func main(){
	fmt.Print(maxProfit([]int{1, 3, 2, 8, 4, 9},2))
}
func max(a int, b int)int{
	if a > b{
		return a
	}else{
		return b
	}
}
func maxProfit(prices []int, fee int) int{
	d := [][2]float64{[2]float64{0.0, -50000.0*50000.0}}
	res := 0.0
	for i:=1;i<len(prices)+1;i++{
		d = append(d, [2]float64{0.0, 0.0})
		d[i][0] = math.Max(d[i-1][0], d[i-1][1]+float64(prices[i-1]-fee))
		d[i][1] = math.Max(d[i-1][1], d[i-1][0]-float64(prices[i-1]))
		if math.Max(d[i][0], d[i][1]) > res{
			res = math.Max(d[i][0], d[i][1])
		}
	}
	return int(res)
}



//func maxProfit(prices []int, fee int) int{
//	d := [][2]int{[2]int{0, math.MinInt64}}
//	res := 0
//	for i:=1;i<len(prices)+1;i++{
//		d = append(d, [2]int{0, 0})
//		d[i][0] = max(d[i-1][0], d[i-1][1]+prices[i-1])
//		d[i][1] = max(d[i-1][1], d[i-1][0]-prices[i-1])
//		if max(d[i][0], d[i][1]) > res{
//			res = max(d[i][0], d[i][1])
//		}
//	}
//	return res
//}

