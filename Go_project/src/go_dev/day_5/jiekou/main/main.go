package main

import (
	"fmt"
)


type Test interface{
	Print()

}

type Student struct{
	name string
	age int
	score int
}

func (p *Student) Print() {
	fmt.Println("name",p.name)
	fmt.Println("age",p.age)
	fmt.Println("score",p.score)
}

func main(){
	
	var t Test
	var stu Student = Student{
		name:"stu1",
		age:20,
		score:200,
	}

	t = &stu //只要一个类型实现了一个接口，那么这个类型就可以赋值给接口变量
	t.Print()

}

