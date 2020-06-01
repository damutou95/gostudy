package main

import (
	"fmt"
)

func main(){
 fmt.Print(isPalindrome(12321))
}
//官方解答的实现
func isPalindrome(x int) bool {
	if x<0 || (x%10==0 && x !=0){
		return false
	}
	rest := 0
	result := 0
	for x>result{
		rest = x % 10
		x /= 10
		result = result * 10 + rest
	}
	return x == result || x == result/10
}
//自己的解法
//func isPalindrome(x int) bool {
//	if x<0{
//		return false
//	}
//	y := x
//	result := 0
//	for x!=0{
//		res := x % 10
//		x /= 10
//		result = result*10 + res
//	}
//	if result == y{
//		return true
//	}
//	return false
//}