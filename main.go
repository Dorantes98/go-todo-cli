package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Todo List:")

	for i, arg := range os.Args { // Print out the args
		fmt.Printf("Arg %d: %s\n", i, arg)
	}

	if os.Args[1] == "add" {
		if len(os.Args) < 3 {		// Checks to make sure the user inputed the correct number of Args
			fmt.Printf("Error: missing task description. ")
			return // Early return to avoid crashes
		}
		fmt.Printf("Adding task: %s\n", os.Args[2]) // Print the description of task being added

		file, err := os.OpenFile("tasks.csv", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
		defer file.Close()

		if err != nil {
			fmt.Println("Error opening file:", err)
		} 

		fmt.Fprintln(file, os.Args[2])
	}

	if os.Args[1] == "list" {
		file, err := os.Open("tasks.csv")
		defer file.Close() 

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}

		err = scanner.Err()
		if err != nil {
			fmt.Println("Error reading file:", err)
		}
	}
}
