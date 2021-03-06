package model

import(
	"errors"
)

var(
	ErrNotFoundBook = errors.New("not found this book...")
)

type Student struct{
	Name string
	Grade string
	Id string
	Sex string
	books []*BorrowItem
}

type BorrowItem struct{
	book *Book
	num int
}

func CreateStudent(name,grade,id,sex string) *Student{
	stu := &Student{
		Name: name
		Grade: grade
		Id: id
		Sex: sex,
	}
	return stu
}

func (s *Student) AddBook(b *BorrowItem){
	s.books = append(s.books,b)
}

func (s *Student) DelBook(b *BorrowItem)(err errors){
	found := false
	for i:= 0; i<len(s.books);i++{
		if(s.books[i].book.Name == b.book.Name){
			if(b.num == s.books[i.num]) {
				right := s.books[0:i]
				left := s.books[i+1:]
				right = append(right,left...)
				s.books =right
				return
			}
			s.books[i].num -= b.num
			return
		}
	}
	err = ErrNotFoundBook
	return
}

func (s *Student) GetBookList()[]*BorrowItem{
	
}
