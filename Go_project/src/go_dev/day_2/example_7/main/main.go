package main

import(
	"fmt"
)

func main(){
	var str string = "hello bob! \n"
	var str1 string = `
	床前明月光，
	疑是地上霜。
	举头望明月，
	低头思故乡。`
	var c byte = 'c' //字符只能是单引号
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(c)
	fmt.Printf("%c\n",c)

	var a int = 100
	var b bool

	fmt.Printf("%v\n",a)
	fmt.Printf("%v\n",b)
	fmt.Printf("%v\n",c)
	fmt.Printf("%T\n",c)
	fmt.Printf("20%%\n")
	fmt.Printf("%b\n",100)
	fmt.Printf("%f\n",100.09)
	fmt.Printf("%q\n","hello")
	fmt.Printf("%p\n",&a)

	str3 := fmt.Sprintf("a=%d",a)
	fmt.Printf("%q\n",str3)
}