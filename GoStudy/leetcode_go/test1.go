package main

import "fmt"
func main(){
s := "hello"
p := "ll"
fmt.Print(strStr(s,p))
}
func getNextMatrix(p string)[]int{
	if len(p) == 0{
		return nil
	}
	next := make([]int,len(p))
	next[0] = -1
	j := 0
	k := -1
	for j<len(p)-1{
		if k == -1 || p[k] == p[j]{
			k++
			j++
			if p[k] != p[j]{
				next[j] = k
			}else{
				next[j] = next[k]
			}
		}else{
			k = next[k]
		}
	}
	return next
}
func strStr(haystack string, needle string) int {
	next := getNextMatrix(needle)
	if len(next) == 0{
		return 0
	}
	i := 0
	j := 0
	for i<len(haystack) && j<len(needle){
		if j== -1 || haystack[i] == needle[j]{
			i++
			j++
		}else{
			j = next[j]
		}
	}
	if j==len(needle){
		return i-j
	}else{
		return -1
	}
}