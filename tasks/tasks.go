package tasks

import (
	"fmt"
	"os"
	"bufio"
)

func AddTasks(description string) error {
	file, err := os.OpenFile("tasks.csv", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("Error opening file: %w", err)
	} 
	defer file.Close()

	_, err = fmt.Fprintln(file, description)
	if err != nil {
		return fmt.Errorf("Failed to write task: %w", err)
	}

	return nil
}

func ListTasks() error {
	file, err := os.Open("tasks.csv")

	if err != nil {
		return fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close() 

	scanner := bufio.NewScanner(file)

	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()
		fmt.Printf("%d. %s\n", i, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: %w", err)
	}

	return nil
}