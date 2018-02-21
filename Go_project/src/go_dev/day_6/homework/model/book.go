package model

import(
	
	"errors"
)

var(
	errStockNotEnough = errors.New("stock is not enough...")
)

type Book struct{
	Name string
	Total int
	Author string
	CreateTime string
}

func CreateBook(name string,total int,author ,createtime string )(b *Book){
	b = &Book{
		Name:name,
		Total:total,
		Author:author,
		CreateTime:createtime
	}
	return
}

func (b *Book) canBorrow(c int) bool{
	return b.Total >= c
}

func (b *Book) Borrow(c int) (err errors){
	if b.canBorrow == false{
		err = errStockNotEnough
		return
	}
	b.Total -= c
	return
}

func (b *Book)Back(c int)(err errors){
	b.Total += c
	return
}
