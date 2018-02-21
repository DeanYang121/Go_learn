package point

import (
	"fmt"
)


func modify(p *int){
	fmt.Println(p)
	*p = 100000
	return
}

func PointExam(){
	var a int =10
	fmt.Println(&a)

	var p *int
	p = &a
	fmt.Println(*p)

	*p = 100
	fmt.Println(a)

	var b int = 999
	p = &b
	*p = 5

	fmt.Println("a",a,"b",b)
	modify(&a)
	fmt.Println(a)

}