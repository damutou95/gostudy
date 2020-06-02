package main

import "fmt"

func main(){
	s :="(){}"
	fmt.Print(isValid(s))
}
func isValid(s string) bool {
	stack := []byte{}
	for _,v:= range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, byte(v))
		} else if v == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}else{
				stack = stack[:len(stack)-1]
				continue
			}

		} else if v == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '['  {
				return false
			}else{
				stack = stack[:len(stack)-1]
				continue
			}

		} else if v == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != '{'{
				return false
			}else{
				stack = stack[:len(stack)-1]
				continue
			}

		}
	}
	if len(stack) == 0{
		return true
	}else{
		return false
	}
}