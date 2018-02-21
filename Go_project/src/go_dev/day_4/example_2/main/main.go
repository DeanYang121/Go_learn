package main

import (
	"fmt"
	"sort"
)



func testMapSort(){
	var a map[int]int
	a = make(map[int]int,5)

	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[19] = 10

	var keys []int
	for k,_ := range a {
		keys = append(keys,k)
	}

	sort.Ints(keys)

	for _,v := range keys{
		fmt.Println(v,a[v])
	}

}


func testMapReverse(){
	var a map[string]int
	a = make(map[string]int,2)
	var b map[int]string
	b = make(map[int]string,2)

	a["a"] = 101
	a["b"] = 102

	for k,v := range a {
		b[v] = k
	}

	fmt.Println(b)
	
}

func main(){
	testMapSort()
	testMapReverse()
}