package main

import (
	"bytes"
	"fmt"
)
func main(){
	a := "11"
	b := "1"
	fmt.Print(addBinary(a,b))
}
func addBinary(a string, b string) string {
	forward := 0
	var length int
	var buffer = bytes.Buffer{}
	if len(a)>len(b){
		length = len(a)
		delta := len(a) - len(b)
		bBuffer := bytes.Buffer{}
		for k:=delta;k>0;k--{
			bBuffer.WriteString("0")
		}
		bBuffer.WriteString(b)
		b = bBuffer.String()

	}else{
		length = len(b)
		delta := len(b) - len(a)
		aBuffer := bytes.Buffer{}
		for k:=delta;k>0;k--{
			aBuffer.WriteString("0")
		}
		aBuffer.WriteString(a)
		a = aBuffer.String()
	}
	for i := length-1;i>=0;i--{
		if a[i] == '1'{
			forward++
		}
		if b[i] == '1'{
			forward++
		}
		if forward%2 == 1{
			buffer.WriteString("1")
		}else{
			buffer.WriteString("0")
		}
		forward/=2
	}
	if forward == 1{
		buffer.WriteString("1")
		length++
	}
	bufferString := []byte(buffer.String())
	m := 0
	n := length - 1
	for m<n{
		bufferString[m],bufferString[n] = bufferString[n],bufferString[m]
		m++
		n--
	}

	return string(bufferString[:])
}
