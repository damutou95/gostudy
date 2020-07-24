package main

import (
	"fmt"
)
func main(){
	fmt.Print(numOfSubarrays([]int{2,2,2,2,5,5,5,8}, 3, 4))
}
//方法一
//func numOfSubarrays(arr []int, k int, threshold int) int {
//	count := 0
//	if len(arr)<k{
//		return 0
//	}
//	for len(arr)>=k {
//		if len(arr) == k && sum(arr) >= k*threshold {
//			count++
//		} else if sum(arr[:k]) >= k*threshold {
//			count++
//		}
//		arr = arr[1:]
//	}
//	return count
//}
//func sum(arr []int)int{
//	res := 0
//	for _,v:= range arr{
//		res += v
//	}
//	return res
//}

func numOfSubarrays(arr []int, k int, threshold int) int {
	sumSetLong := []int{arr[0]}
	count := 0
	value := k*threshold
	for i:=1;i<len(arr);i++{
		sumSetLong = append(sumSetLong, arr[i]+sumSetLong[i-1])
	}
	for i:=k-1;i<len(arr);i++{
		if i==k-1{
			if sumSetLong[i]>=value{
				count++
			}
		}else if sumSetLong[i]-sumSetLong[i-k]>=value{
			count++
		}

	}
	return count
}

