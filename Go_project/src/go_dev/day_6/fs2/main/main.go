package main

import(
	"reflect"
	"fmt"
	"encoding/json"
)

type Student struct{
	Name string `json:"stu_name"`
	Age int
	Score float32
}

func(s Student) Set(name string,age int,score float32){
	s.Name = name
	s.Age = age
	s.Score = score
}

func(s Student) Print(){
	fmt.Println("hello",s)
}

func TestStruct(a interface{}){
	val := reflect.ValueOf(a)
	tap := reflect.TypeOf(a)
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct!!")
		return
	}

	num := val.Elem().NumField()

	tag := tap.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag is %s\n",tag)
	
	val.Elem().Field(0).SetString("niu pi a")

	for i:= 0;i<num;i++{
		fmt.Printf("%d %v\n",i,val.Elem().Field(i).Kind())
	}

	fmt.Printf("struct has %d fields\n",num) //获取结构体字段数量

	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d methods\n",numOfMethod)
	var paras []reflect.Value
	val.Elem().Method(0).Call(paras)

}

func main(){
	var a Student = Student{
		Name:"dean",
		Age:21,
		Score:80.3,
	}

	result,_ := json.Marshal(a)
	fmt.Println("json result:",string(result))

	TestStruct(&a)
	fmt.Print(a)
}