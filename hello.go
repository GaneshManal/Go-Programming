package main

import (
	"fmt"

	"rsc.io/quote"
)

var id int = 15
var name string = "Neha"
var age float32 = 22

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}
