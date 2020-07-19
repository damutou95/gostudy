package main
import "fmt"
func main(){
	fmt.Print(maxProfit([]int{7,1,5,3,6,4}))
}

//官方法二
//func maxProfit(prices []int) int {
//	maxValue := 0
//	peak := prices[0]
//	valley := prices[0]
//	i := 0
//	for i<len(prices)-1{
//		for i<len(prices)-1 && prices[i]>prices[i+1]{
//			i++
//		}
//		valley = prices[i]
//		for i<len(prices)-1 && prices[i]<prices[i+1]{
//			i++
//		}
//		peak = prices[i]
//		maxValue += peak - valley
//	}
//	return maxValue
//}

//官方法三
func maxProfit(prices []int) int {
	maxValue := 0
	for i:=0;i<len(prices)-1;i++{
		if prices[i+1]>prices[i]{
			maxValue += prices[i+1] - prices[i]
		}else{
			continue
		}
	}
	return maxValue
}