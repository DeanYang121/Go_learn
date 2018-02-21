package main

import(
	_ "fmt"
	_ "os"
	"go_dev/day_7/IOop/buffIo"
)

// type Student struct{
// 	Name string
// 	Age int
// 	Score float32
// }

func main(){
// 	// fmt.Fprintf(os.Stdout,"hello world")
// 	// file,err:= os.OpenFile("./test.log",os.O_CREATE|os.O_WRONLY,0664)
// 	// if err != nil{
// 	// 	fmt.Println("open file err,",err)
// 	// 	return
// 	// }
// 	// fmt.Fprintf(file,"hello world hahhaa")
// 	// file.Close()

// 	var str = "stu01 20 300"
// 	var stu Student
// 	fmt.Sscanf(str,"%s %d %f",&stu.Name,&stu.Age,&stu.Score) //格式化字符串输入
// 	fmt.Println(stu)

	buffIo.TestBufio2()

}

