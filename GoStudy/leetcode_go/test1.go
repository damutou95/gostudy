package main
import "fmt"
func main(){
	a := [][]int{[]int{1,2,3}}
	a = append([][]int{[]int{}},a...)
	fmt.Print(a)
}