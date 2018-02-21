package main

import(
	"fmt"
)

type Student struct{
	Number int
}

type stu Student  //类型别名



func main(){
	var a Student
	a = Student{30}

	var b stu
	a = Student(b)  //虽然是别名 但是还是需要强制转换
	fmt.Println(a) 
}