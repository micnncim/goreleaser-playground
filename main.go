package main

import (
	"fmt"

	"rsc.io/quote"
)

const version = "0.2.7"

func main() {
	fmt.Println("version: " + version)
	fmt.Println(quote.Hello())
}
