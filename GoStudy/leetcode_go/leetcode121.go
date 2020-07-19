package main

import "fmt"
func main(){
	fmt.Print(maxProfit([]int{7,6,1,4,3,9,1}))
}
//递归
//func max(x int, y int, z int)int{
//	if x > y && x > z{
//		return x
//	}else if y > x && y > z{
//		return y
//	}else{
//		return z
//	}
//}
//func maxHelper(prices []int) int {
//	if len(prices) < 2{
//		return 0
//	}
//	res := prices[len(prices)-1] - prices[0]
//	for i:=1;i<len(prices);i++{
//		if prices[len(prices)-1] - prices[i] > res{
//			res = prices[len(prices)-1] - prices[i]
//		}
//	}
//	return res
//}
//func minHelper(prices []int) int {
//	if len(prices) < 2{
//		return 0
//	}
//	res := prices[1] - prices[0]
//	for i:=2;i<len(prices);i++ {
//		if prices[i]-prices[0] > res {
//			res = prices[i] - prices[0]
//		}
//	}
//	return res
//}
//func maxProfit(prices []int) int {
//	if len(prices) < 2{
//		return 0
//	}
//	minIndex := 0
//	maxIndex := 0
//	for i:=0;i<len(prices);i++{
//		if prices[i]>prices[maxIndex]{
//			maxIndex = i
//		}
//		if prices[i]<prices[minIndex]{
//			minIndex = i
//		}
//	}
//	if maxIndex >= minIndex{
//		return prices[maxIndex] - prices[minIndex]
//	}else{
//		return max(maxHelper(prices[:maxIndex+1]),maxProfit(prices[maxIndex+1:minIndex]),minHelper(prices[minIndex:]))
//	}
//
//}


//官方题解
//func maxProfit(prices []int) int {
//	if len(prices) < 2{
//		return 0
//	}
//	minValueSet := []int{prices[0]}
//	for i:=1;i<len(prices);i++{
//		if prices[i] < minValueSet[i-1]{
//			minValueSet = append(minValueSet, prices[i])
//		}else{
//			minValueSet = append(minValueSet, minValueSet[i-1])
//		}
//	}
//	maxDelta := 0
//	for j:=1;j<len(prices);j++{
//		if maxDelta < prices[j] - minValueSet[j]{
//			maxDelta = prices[j] - minValueSet[j]
//		}
//	}
//	return maxDelta
//}
//改良版本，把队列去掉，只用一个变量保存到目前为止的最大利润。
func maxProfit(prices []int) int {
	if len(prices) < 2{
		return 0
	}
	minValue := prices[0]
	maxDelta := 0
	for i:=1;i<len(prices);i++{
		if prices[i] < minValue{
			minValue = prices[i]
		}
		if maxDelta < prices[i] - minValue{
			maxDelta = prices[i] - minValue
		}
	}
	return maxDelta
}