package main

import (
	"fmt"

	"rsc.io/quote"
)

const version = "0.1.0"

func main() {
	fmt.Println("version: " + version)
	fmt.Println(quote.Hello())
}
