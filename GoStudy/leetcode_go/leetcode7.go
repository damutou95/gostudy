package main

import (
	"fmt"
)

func main(){
	var x int
	x = -120300
	fmt.Print(reverse(x))
}

func reverse(x int) int {
	res := 0
	last := 0
	for x != 0{
		last = x % 10
		x = x / 10
		if res > 214748364 || (res == 214748364 && last>7){
			return 0
		}
		if res < -214748364 || (res == -214748364 && last<(-8)){
			return 0
		}
		res = res*10 + last
	}
	return res
}
