package work

import(
	"fmt"
)

func Shui(){

	for i :=100 ;i<1000;i++{
		a := int(i/100)
		a_1 := i%100
		b := int(a_1/10)
		c := int(a_1%10)
		d := a*a*a + b*b*b+c*c*c
		if( i ==  d ){
			fmt.Println("这是一个水仙花数：",i)
		}

	}
}