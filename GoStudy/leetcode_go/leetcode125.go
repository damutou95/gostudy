package main
import (
	"fmt"
	"strings"
	"unicode"
)
func main(){
	fmt.Print(isPalindrome("A man  plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
	head := 0
	tail := len(s) - 1
	for head < tail{
		if head < tail && !unicode.IsDigit(rune(s[head])) && !unicode.IsLetter(rune(s[head])){
			head++
			continue
		}
		if head < tail && !unicode.IsDigit(rune(s[tail])) && !unicode.IsLetter(rune(s[tail])){
			tail--
			continue
		}
		if strings.ToLower(string(s[head])) == strings.ToLower(string(s[tail])){
			head++
			tail--
			continue
		}else{
			return false
		}

	}
	return true
}