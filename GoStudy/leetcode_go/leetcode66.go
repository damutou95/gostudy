package main
import "fmt"
func main(){
	digits := []int{9,9,9}
	fmt.Print(plusOne(digits))
}
func plusOne(digits []int) []int {
	forward := 1
	for i := len(digits)-1;i>=0;i--{
		if digits[i] == 9 && forward == 1{
			forward = 1
			digits[i] = 0
		}else{
			digits[i] = digits[i] + forward
			forward = 0
		}
	}
	if forward == 0{
		return digits
	}else{
		 digits = append([]int{1,},digits...)
	}
	return digits
}
