package main

import (
	"fmt"
)


type Reader interface{
	Read()
}

type Writer interface{
	Write()
}

type ReadWrite interface{
	Reader
	Writer
}

type File struct{

}

func (f *File) Read(){
	fmt.Println("read data")
}

func (f *File) Write(){
	fmt.Println("write data")
}

func Test(rw ReadWrite){
	rw.Read()
	rw.Write()
}

func main(){
	var f *File //类型要一致
	var b interface{}
	b = f
	// Test(&f)
	fmt.Println("111...")
	if v,ok := b.(ReadWrite);ok{
		fmt.Println(v,ok)
	}

}