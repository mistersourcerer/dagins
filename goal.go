package main

import (
	"fmt"
)

func HandleGoal(cmd []string) {
	switch command := cmd[0]; command {
	case "create":
		// handle abscence of a name...
		create(cmd[1])
	}
}

func create(name string) {
	fmt.Println("Creating", name)
}
