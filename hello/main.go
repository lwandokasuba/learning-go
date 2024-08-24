package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello Lwando!")
	fmt.Println(("Some quotes for you:"))
	fmt.Println("Go quote:", quote.Go())
	fmt.Println("Hello quote:", quote.Hello())
	fmt.Println("Glass quote:", quote.Glass())
	fmt.Println("Optimization quote:", quote.Opt())
}
