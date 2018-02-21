package main

import(
	"fmt"
	"time"
)

const(
	man = iota+1
	female 
)

func main(){
	fmt.Println(man,female)
	//1970年到现在的秒数  查官方文档的包--》index--》查函数
	for{
		_time_now := time.Now().Unix()
		if(_time_now%female == 0){
			fmt.Println("female =",female)
		}else{
			fmt.Println("man =",man)
		}
		time.Sleep(1000*time.Microsecond)
	}

}