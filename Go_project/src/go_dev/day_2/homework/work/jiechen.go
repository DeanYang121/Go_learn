package work

import (
	"fmt"
)


func Jiecheng(){
	n := 5
	sum := 0
	for i:=1;i<=n;i++{
		sum += jie(i)
	}
	fmt.Printf("%d阶乘之和为=%d",n,sum)
}

func jie(n int) int{
	mul := 1
	for i := 1;i<=n;i++{
		mul *= i
	}
	return mul
}