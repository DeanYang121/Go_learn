package test

import(
	"fmt"
)

var Name string = "dean!"
var Age int = 2100

func init(){
	fmt.Println("this is test")
	fmt.Println("test.package.Name =",Name)
	Age = 10
	fmt.Println("test.package.Age =",Age)
}