package main

import "fmt"

type Book struct {
	Name   string
	Author string
	Price  float32
}

func main() {
	b := Book{Name: "Chess", Author: "Stefan Zweig", Price: 9.50}
	fmt.Printf("Book data: %+v \n", b)
	fmt.Printf("Book data: %v \n", b)
	fmt.Printf("Name: %s \n", b.Name)
	fmt.Printf("Author: %s \n", b.Author)
	fmt.Printf("Price: %f \n", b.Price)

	/*
	   Book data: {Name:Chess Author:Stefan Zweig Price:9.5}
	   Book data: {Chess Stefan Zweig 9.5}
	   Name: Chess
	   Author: Stefan Zweig
	   Price: 9.500000
	*/
}
