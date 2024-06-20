package main

import (
	"flag"
	"fmt"
	// "flag"
)

func main() {
	var findRecursive = flag.Bool("r", false, "what is this")
	flag.Parse()
	fmt.Print("hej", *findRecursive)
}
