package main

import (
	"fmt"
)

func reverse(str string) string{
	var result string
	strlen := len(str)

	for i:=0;i<strlen;i++{
		result = result+fmt.Sprintf("%c",str[strlen-i-1])
	}
	return result
}
func reverse_1(str string)string{
	var result []byte
	tmp := []byte(str)
	length := len(tmp)
	for i := 0;i<length;i++{
		result = append(result,tmp[length-i-1])
	}
	return string(result)
}

func main(){
	var str1 string = "hello"
	str2 := "world"
	//str3 := str1 +" "+str2

	str3 := fmt.Sprintf("%s %s",str1,str2)
	n := len(str3)

	fmt.Println(str3)
	fmt.Printf("len(str3)=%d\n",n)

	sub_str := str3[0:5]
	fmt.Println(sub_str)

	sub_str = str3[6:]
	fmt.Println(sub_str)

	fmt.Println(reverse(str3))
	fmt.Println(reverse_1(str3))
}