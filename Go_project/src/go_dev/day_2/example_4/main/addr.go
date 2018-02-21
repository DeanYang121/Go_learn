package main

import(
	"fmt"
)
//指针可以修改值类型
func modify(a *int){
	*a = 10
}

func modify_1(a int){
	a =20
}

func main(){
	a := 5
	b := make(chan int,1)

	fmt.Println("a=",a)
	fmt.Println("b=",b)

	modify(&a)
	fmt.Println("a=",a)

	modify_1(a)
	fmt.Println("a=",a)
}