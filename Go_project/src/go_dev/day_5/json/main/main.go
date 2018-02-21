package main

import (
	"time"
	"fmt"
	_ "encoding/json"
)

type Student struct{
	Name string `json:"student_name"`  //json tag标签 告诉json里面的key是什么
	Age int		`json:"age"`	
	Score int   `json:"score"`
}

type Cart struct{
	name string
	age int
}

type Train struct{
	Cart
	int
	start time.Time
	age int
}

func main(){
	// var stu Student = Student{
	// 	Name:"stu01",
	// 	Age:18,
	// 	Score:80,
	// }

	// data, err := json.Marshal(stu)
	// if err != nil{
	// 	fmt.Println("json encode stu failed...error:",err)
	// }

	// fmt.Println(string(data))
	var t Train
	t.name = "train" //相当于 t.Cart.age = 200
	t.age = 100     //相当于 t.Cart.name = "Train"
	t.int = 200

	fmt.Println(t)


}