package main


import(
	"go_dev/day_1/goroute_example/goroute"
	"fmt"
)

func main(){
	var pipe chan int
	pipe = make(chan int,1)
	go goroute.Add(200,300,pipe)
	sum := <- pipe
	fmt.Println("sum =",sum)
}