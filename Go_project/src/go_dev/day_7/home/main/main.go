package main

import(
	"go_dev/day_7/home/balance"
	"fmt"
	"math/rand"
	"time"
	"os"
)

func main(){

	var insts []*balance.Instance
	for i := 0;i < 5;i++{
		host := fmt.Sprintf("192.168.%d.%d",rand.Intn(255),rand.Intn(255))
		one := balance.NewInstance(host,8080)
		insts = append(insts,one)
	}

	// var balancer balance.Balance
	// var conf = "random"
	var balanceName = "random"
	if len(os.Args) > 1{
		balanceName = os.Args[1]
	}

	// if conf == "random"{
	// 	balancer = &balance.RandomBalance{}
	// 	fmt.Println("use random balancer")
	// }else if conf == "roundrobin" {
	// 	balancer = &balance.RoundRobinBalance{} 
	// 	fmt.Println("use roundrobin balancer")
	// }


	for{
		inst,err := balance.DoBalance(balanceName,insts)
		if err != nil{
			fmt.Println("do balance err:",err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}


}