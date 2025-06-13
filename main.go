package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Dorantes98/go-todo-cli/tasks"
)

func main() {
	// for i, arg := range os.Args { // Print out the args
	// 	fmt.Printf("Arg %d: %s\n", i, arg)
	// }

	if os.Args[1] == "add" {
		if len(os.Args) < 3 { // Checks to make sure the user inputed the correct number of Args
			fmt.Printf("Error: missing task description. ")
			return // Early return to avoid crashes
		}

		desc := os.Args[2]
		fmt.Printf("Adding task: %s\n", desc)

		err := tasks.AddTasks(desc)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	if os.Args[1] == "list" {
		err := tasks.ListTasks()
		if err != nil {
			fmt.Println("Error:", err)
		}
		return
	}

	if os.Args[1] == "complete" {
		if len(os.Args) < 3 {
			fmt.Println("Missing task ID.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		err := tasks.CompleteTask(id)
		if err != nil {
			fmt.Println("Error:", err)
		}
		return
	}
}
