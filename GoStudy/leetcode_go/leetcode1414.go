package main
import "fmt"
func main(){
	fmt.Print(findMinFibonacciNumbers(5))
}

func findMinFibonacciNumbers(k int) int {
	fibonacci := []int{1,1}
	a := 1
	b := 1
	for b < k{
		b += a
		a = b - a
		fibonacci = append(fibonacci, b)
	}
	count := 0
	for j:=len(fibonacci)-1;j>=0;j--{
		if fibonacci[j] > k{
			continue
		}else{
			k -= fibonacci[j]
			count++
			if k == 0{
				return count
			}
		}
	}
	return count
}
