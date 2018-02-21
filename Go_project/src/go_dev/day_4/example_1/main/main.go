package main


import(
	"fmt"
	"time"
	"errors"
)

func initConfig()(err error){
	return errors.New("init config failed")
}
//20

func test(){

	// defer func(){
	// 	if err := recover();err != nil{
	// 		fmt.Println(err)
	// 	}
	// }()

	err := initConfig()
	if err != nil{
		panic(err)
	}
	return 
	// b := 0
	// a := 100 / b
	fmt.Println("a")
	// return
}

func main(){
	// var i int
	// fmt.Println(i)

	// j := new(int)
	// *j = 100
	// fmt.Println(*j)
	// var a []int //一个slice
	// a = append(a,10,20,383)
	// a = append(a,a...) //添加一个slice  ...可以表示可变长参数
	// fmt.Println(a)
	for{
		test()
		time.Sleep(time.Second)
	}


}