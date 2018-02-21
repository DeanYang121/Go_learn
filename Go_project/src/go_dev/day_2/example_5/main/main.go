package main

import(
	"fmt"
)

func swap(a *int,b *int){
	tmp := *a
	*a = *b
	*b = tmp
	return
}

func main(){
	first := 100
	second := 200
	fmt.Println("first:",first,"second:",second)	
	swap(&first,&second)
	fmt.Println("after swap...")
	fmt.Println("first:",first,"second:",second)
}