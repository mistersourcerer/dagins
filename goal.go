package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type goal struct {
	Name        string
	Periodicity string
}

func HandleGoal(cmd []string) {
	switch command := cmd[0]; command {
	case "create":
		if len(cmd[1:]) < 1 {
			fmt.Println("goal need a name")
			fmt.Println("  $ dagins goal create 'read one page'")
		} else {
			// extract periodicity
			create(cmd[1], "daily")
		}
	}
}

func create(name string, periodicity string) {
	// validate periodicity
	fmt.Println("Creating", name, periodicity)
	newGoal := goal{Name: name, Periodicity: periodicity}
	save(&newGoal)
}

func save(newGoal *goal) {
	// handle error
	os.Mkdir(".dagins", 0755)

	f, err := os.OpenFile(".dagins/goals.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	// handle error
	goalJSON, err := json.Marshal(newGoal)

	// handle error
	f.Write(goalJSON)
	f.Write([]byte("\n"))
}
