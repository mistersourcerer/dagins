package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type goal struct {
	Name        string
	Periodicity string
}

func HandleGoal(cmd []string) {
	switch command := cmd[0]; command {
	case "create":
		create(cmd)
	}
}

func create(parameters []string) {
	if len(parameters[1:]) < 1 {
		fmt.Println("goal need a name")
		fmt.Println("  $ dagins goal create 'read one page'")
	} else {
		name := parameters[1]

		// extract periodicity
		periodicity := "daily"

		// validate periodicity
		fmt.Println("Creating", name, periodicity)
		newGoal := goal{Name: name, Periodicity: periodicity}
		save(&newGoal)
	}
}

func save(newGoal *goal) {
	goalJSON, err := json.Marshal(newGoal)
	if err != nil {
		panic(err)
	}
	fileAppend(goalJSON)
}

func fileAppend(goalJSON []byte) {
	path := ".dagins"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}

	f, err := os.OpenFile(".dagins/goals.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// handle error
	f.Write(goalJSON)
	f.Write([]byte("\n"))
}
