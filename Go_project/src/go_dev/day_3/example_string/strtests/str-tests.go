package strtests

import (
	"fmt"
	"strings"
	"strconv"
)

func Strlearn(){
	
	str := "hello world abc \n"
	result := strings.Replace(str,"world","you",1)
	fmt.Println(result)

	count := strings.Count(str,"l") //计算l出现的次数
	fmt.Println(count)

	result = strings.Repeat(str,3) //重复三次打印出来
	fmt.Println(result)

	result = strings.ToLower(str)
	fmt.Println(result)

	result = strings.ToUpper(str)
	fmt.Println(result)

	result = strings.TrimSpace(str) //去除两边的空白
	fmt.Println(result)

	result = strings.Trim(str,"\n\r")
	fmt.Println(result)

	result = strings.TrimLeft(str,"\n\r")
	fmt.Println(result)

	result = strings.TrimRight(str,"\n\r")
	fmt.Println(result)

	splitResult := strings.Fields(str)//指定分隔符
	for i:=0;i<len(splitResult);i++{
		fmt.Println(splitResult[i])
	}

	splitResult = strings.Split(str,"l")
	for i:=0;i<len(splitResult);i++{
		fmt.Println(splitResult[i])
	}

	str2 := strings.Join(splitResult,"l")
	fmt.Println("join",str2)

	str2 =strconv.Itoa(100)
	fmt.Println(str2)

	number,err := strconv.Atoi(str2)
	if err != nil{
		fmt.Println("can not convert to int ",err)
	}else{
		fmt.Println(number)
	}

}