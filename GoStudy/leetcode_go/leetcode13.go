package main

import "fmt"

func main(){
	  fmt.Print(romanToInt("XXIV"))
}
func romanToInt(s string) int {
	dict := map[byte]int{
		'I':1,
		'V':5,
		'X':10,
		'L':50,
		'C':100,
		'D':500,
		'M':1000,
	}
	result := 0
	temp := 1000
	for _,char:=range s{
		if dict[byte(char)]>temp{
			result = result - temp + dict[byte(char)] - temp
		}else{
			result = result + dict[byte(char)]
		}
		temp = dict[byte(char)]
	}
	return result
}