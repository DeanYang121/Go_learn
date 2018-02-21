package control

import (
	"strconv"
	"fmt"
)


func Control(){
	var str string
	fmt.Scanf("%s",&str)

	number,err := strconv.Atoi(str)
	if err != nil{
		fmt.Println("can not convert str to int")
	}else{
		fmt.Println(number)
	}

}