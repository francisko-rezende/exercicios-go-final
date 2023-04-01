package main

import "fmt"

type Book struct {
	Title         string
	Author        string
	PublishedYear int
	Pages         int
}

func main() {
	b := Book{"A Guerra dos Mundos", "H. G. Wells", 1898, 303}
	fmt.Println(b)
}
