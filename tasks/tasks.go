package tasks

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	IsComplete  bool
}

func AddTasks(description string) error {
	file, err := os.OpenFile("tasks.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Could not get file info: %w", err)
	}

	isNewFile := stat.Size() == 0
	file.Seek(0, 0)
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	recordCount := len(records)
	if !isNewFile {
		recordCount--
	}
	nextID := recordCount + 1

	task := Task{
		ID:          nextID,
		Description: description,
		CreatedAt:   time.Now(),
		IsComplete:  false,
	}

	writer := csv.NewWriter(file)
	if isNewFile {
		writer.Write([]string{"ID", "Description", "CreatedAt", "IsComplete"})
	}

	record := []string{
		strconv.Itoa(task.ID),
		task.Description,
		task.CreatedAt.Format(time.RFC3339),
		strconv.FormatBool(task.IsComplete),
	}

	if err := writer.Write(record); err != nil {
		return fmt.Errorf("Failed to write task: %w", err)
	}

	writer.Flush()
	return nil
}

func ListTasks() error {
	file, err := os.Open("tasks.csv")

	if err != nil {
		return fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()

	fmt.Println("\nTodo List:")

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
