package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	version := flag.Bool("version", false, "Show this dagins version.")
	v := flag.Bool("v", false, "Show this dagins version.")
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()

		return
	}

	goal := flag.NewFlagSet("goal", flag.ExitOnError)
	goal.Parse(os.Args[2:])

	switch {
	case *version || *v:
		showVersion()
	default:
		// currently only handle goal subcommand. :(((
		// haha, it is fine though. Still very much just exploring :)
		if len(os.Args[2:]) == 0 {
			goal.Usage()
		} else {
			HandleGoal(os.Args[2:])
		}
	}
}

func showVersion() {
	fmt.Println("dagins version 0.0.1a")
}
