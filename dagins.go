package main

import (
	"flag"
	"fmt"
)

func main() {
	version := flag.Bool("version", false, "Show this dagins version.")
	v := flag.Bool("v", false, "Show this dagins version.")
	flag.Parse()

	switch {
	case *version || *v:
		showVersion()
	default:
		dagins()
	}
}

func showVersion() {
	fmt.Println("dagins version 0.0.1a")
}

func dagins() {
	fmt.Println("JUST DO IT! STAY ON TARGET! YOURS TRULY, DAGINS!")
}
