package main
import(
	"reflect"
	"fmt"
)

type Student struct{
	Name string
	Age int
	Score float32
}

func test(b interface{}){
	t := reflect.TypeOf(b)
	fmt.Println(t)

	v := reflect.ValueOf(b)
	k := v.Kind() //kind只是类别 比如struct只显示struct
	fmt.Println(k)

	iv := v.Interface()
	stu,ok := iv.(Student)
	if ok{
		fmt.Printf("%v %T\n",stu,stu)
	}
}

func testInt(b interface{}){
	val := reflect.ValueOf(b)
	val.Elem().SetInt(100) //elem作用是起一个*的作用，代表给指针赋值

	c := val.Elem().Int()
	fmt.Printf("get value of interface{} %d\n",c)
	fmt.Printf("string val:%d\n",val.Elem().Int())
}

func main(){
	var a int = 200
	test(a)
	var b Student = Student{
		Name: "deanyang",
		Age: 32,
		Score:92,
	}
	test(b)
	// testInt(1234)

	var c int = 1
	c = 200
	testInt(&c)
	fmt.Println(c)

}