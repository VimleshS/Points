package main

import (
	"fmt"
)

type MyHandlerFunc func(int)

func (h MyHandlerFunc) PrintMe() {
	fmt.Printf("From Handler.. %d", h)
}

func main() {
	fmt.Println("Starats")
	fn := func(i int) {
		fmt.Println("in funct %d \n", i)
	}

	fn(10)
	f := MyHandlerFunc(fn) // HandlerFunc
	f(23
	f.PrintMe()
}
