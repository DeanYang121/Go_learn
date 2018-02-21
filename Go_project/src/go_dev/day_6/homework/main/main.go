package main

import (
	"fmt"
)

type Car struct{
	Name string
	Age int
}

type Train struct{
	Car
	int
}

func (train *Train) String() string{
	str := fmt.Sprintf("name=[%s] age=[%d] count=[%d]",train.Name,train.Age,train.int)
	return str
}

func main(){
	var train Train
	train.int = 300
	train.Name = "test"
	fmt.Println(&train)
}

