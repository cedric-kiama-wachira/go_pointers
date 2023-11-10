package main

import (
	"fmt"
)

func main() {
	s := "hello"
	fmt.Printf("%T %v \n", s, s)
	ps := &s
	*ps = "word"
	fmt.Printf("%T %v \n", s, s)
}
