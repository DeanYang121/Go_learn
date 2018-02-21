package main

import (
	"fmt"
)


type Car struct{
	weight int
	name string
}

func (p *Car) Run(){
	fmt.Println("running...")
}

type Bike struct{
	Car
	wheel int
}

type Train struct{
	c Car
} //组合 必须有c Car

//这里不用指正p *Train  或者 &b 才能使用指正
func (p *Train) String()string{
	str := fmt.Sprintf("name=[%s] weight=[%d]",p.c.name,p.c.weight)
	return str
}

func main(){
	var a Bike
	a.name = "bike1"
	a.weight = 10
	a.wheel = 2

	fmt.Println(a)
	a.Run()

	var b Train
	b.c.weight = 100
	b.c.name = "Train-1"
	fmt.Println(&b)

}