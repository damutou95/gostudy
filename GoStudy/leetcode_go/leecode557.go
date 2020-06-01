package main
import "fmt"
func main(){
	s := "hello world"
	fmt.Print(reverseWords(s))

}
func reverseWords(s string) string{
	b := []rune(s)
	l := 0
	for i,v := range s{
		if i == len(s)-1 || v == ' '{
			r := i - 1
			if i == len(s)-1{
				r = len(s)-1
			}
			for l < r{
				b[l],b[r] = b[r],b[l]
				l++
				r--
			}
			l = i + 1
		}
	}
	return string(b)
}