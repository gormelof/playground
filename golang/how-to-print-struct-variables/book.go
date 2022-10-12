package playground

import "fmt"

type Book struct {
	Name   string
	Author string
	Price  float32
}

func (b *Book) print() {
	book := Book{Name: "Chess", Author: "Stefan Zweig", Price: 9.50}
	fmt.Printf("Book data: %+v \n", book)
	fmt.Printf("Book data: %v \n", book)
	fmt.Printf("Name: %s \n", book.Name)
	fmt.Printf("Author: %s \n", book.Author)
	fmt.Printf("Price: %f \n", book.Price)

	/*
	   Book data: {Name:Chess Author:Stefan Zweig Price:9.5}
	   Book data: {Chess Stefan Zweig 9.5}
	   Name: Chess
	   Author: Stefan Zweig
	   Price: 9.500000
	*/
}
