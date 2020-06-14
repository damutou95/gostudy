package main

import (
	"bytes"
	"fmt"
	"strconv"
)
func main(){
fmt.Print(countAndSay(4))
}
func countAndSay(n int) string {
	if n == 1{
		return "1"
	}
	s := countAndSay(n-1)
	lastLetter := s[0]
	count := 1
	var res bytes.Buffer
	for _,v:= range s[1:]{
		if byte(v) != lastLetter{
			res.WriteString(strconv.Itoa(count)+string(lastLetter))
			lastLetter = byte(v)
			count = 1
		}else{
			count++
		}
	}
	res.WriteString(strconv.Itoa(count)+string(lastLetter))
	return res.String()
}