package main
import "fmt"
func main() {
	var s string
	s = "abcda"
	res := [] bool{}
	queries := [5][3]int{
		{3,3,0},
		{1,2,0},
		{0,3,1},
		{0,3,2},
		{0,4,1}}
	fmt.Printf("数组的长度是%v\n",len(queries))
	for i:=0;i<len(queries);i++ {
		//fmt.Println("结果是",huiwen(s, queries[i]))
		res = append(res, huiwen(s, queries[i]))
}
	fmt.Print(res)
}


func huiwen(s string, m[3] int) bool{
	var count int = 0
	var temp string
	temp = ""
	for n:=m[0];n<=m[1];n++{
		temp += string(s[n])
	}

	top := m[2]
	var dict map[byte] int
	dict = make(map[byte]int)
	for i:=0;i<len(temp);i++ {
		if dict == nil {


			dict[temp[i]] = 1
		} else {
			value, ok := dict[temp[i]]
			if !ok {
				dict[temp[i]] = 1
			} else {
				dict[temp[i]] = value ^ 1
			}
		}
	}

	for _,v := range dict{

		if v%2 == 1{
			count += 1
		}
	}
	if count/2 <= top{
		return true
	}else{
		return false
	}
}