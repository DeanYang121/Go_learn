package main

import (
	"fmt"
	_ "fmt"
)

type Car interface{
	GetName() string
	Run()
	DiDi()
}

type Test interface{
	Hello()
}

type BMW struct{
	Name string
}

type BYD struct{
	Name string
}

func (bmw *BMW) GetName()string{
	return bmw.Name
}

func (bmw *BMW) Run(){
	fmt.Printf("%s running...\n",bmw.Name)
}

func (bmw *BMW) DiDi(){
	fmt.Println("%s di di di....\n",bmw.Name)
}

func (bmw *BMW) Hello(){
	fmt.Println("hello..",bmw.Name)
}

func main(){
	// var a interface{} //空接口可以保存任何类型的变量
	// var b int
	// a = b
	// var c string
	// a = c
	// fmt.Printf("type of a %T\n",a)
	var car Car
	var bmw BMW
	bmw1 := &BMW{
		Name:"A6",
	}
	bmw.Name = "A5"
	car = &bmw
	car.Run()
	car = bmw1
	car.Run()

	// byd := &BYD{
	// 	Name:"BYD-A7",
	// }
	// car = byd
	var test Test
	test = &bmw
	test.Hello()


}