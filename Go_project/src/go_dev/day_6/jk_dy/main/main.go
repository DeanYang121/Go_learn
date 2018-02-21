package main

import (
	"fmt"
)

type Student struct{
	Name string
	Sex string
}

func Test(a interface{}){
	b,ok := a.(Student)
	if ok == false {
		fmt.Println("convert failed")
		return
	}
	fmt.Println(b)
}

func just(items ...interface{}){
	for index,v :=range items{
		switch v.(type){
		case bool:
			fmt.Printf("%d params is %T,%v\n",index,v,v)
		case int,int64,int32:
			fmt.Printf("%d params is %T,%v\n",index,v,v)
		case float32,float64:
			fmt.Printf("%d params is %T,%v\n",index,v,v)
		case string:
			fmt.Printf("%d params is %T,%v\n",index,v,v)
		case Student:
			fmt.Printf("%d params is %T,%v\n",index,v,v)
		}
	}

}

func main(){
	// var a interface{}
	var b Student = Student{
		Name:"hello",
		Sex:"iii",
	}
	Test(b)
	just(24,4.09,"thisis")
	just(b)
	// a = b
	// c := a.(int)
	// fmt.Printf("%d %T\n",a,a)
	// fmt.Printf("%d %T\n",c,c)
}