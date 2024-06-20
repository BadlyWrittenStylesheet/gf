package main

import (
	"flag"
	"fmt"
	// "flag"
)

func main() {
	var findRecursive = flag.Bool("r", false, "what si ziz")
	flag.Parse()
	fmt.Print("hej", *findRecursive)
}
