package main
import "fmt"
func main(){
	n := 2
	fmt.Print(climbStairs(n))
}
func climbStairs(n int) int {
	a := 0
	b := 1
	for n>0{
		b = a + b
		a = b - a
		n--
	}
	return b
}
