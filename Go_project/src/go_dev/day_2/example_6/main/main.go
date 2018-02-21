package main

import(
	"math/rand"
	"fmt"
	"time"
)

func init(){
	rand.Seed(time.Now().Unix())
}

func main(){

	for i:=0;i<10;i++{
		num := rand.Int()
		fmt.Println(num)
	}

	for i:=0 ; i <10;i++{
		a := rand.Intn(100)
		fmt.Println(a)
	}

	for i:= 0;i <10;i++{
		a := rand.Float32()
		fmt.Println(a)
	}

}



