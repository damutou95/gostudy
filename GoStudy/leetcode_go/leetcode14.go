package main
import "fmt"
func main(){
    var s []string
    s = []string{"aa","a"}
    fmt.Print(longestCommonPrefix(s))
}
func longestCommonPrefix(strs []string) string {
	indexNumber := 0
	if len(strs) == 0{
		return ""
	}
	firstString := strs[0]
	for indexNumber < len(firstString) {
		for _, v := range strs[1:] {
			if len(v)<indexNumber+1 || v[indexNumber] != firstString[indexNumber] {
				return firstString[:indexNumber]
			}
		}
		indexNumber++
	}
	return firstString[:indexNumber]
}