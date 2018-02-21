package work

import (
	"fmt"
)

func Sushu(){
	var result []int
	for i:= 101;i<=200;i++{
		for j := 2;j<=200;j++{
			if(i%j==0 && i!=j){
				break
			}else{
				if(j == 200){
					result = append(result,i)
				}
			}
		}
	}

	fmt.Println("100-200的素数个数：",len(result),result)
}