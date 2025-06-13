package tasks

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
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

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("Error reading CSV: %w", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tDescription\tCreated\tDone")

	fmt.Println("\nTodo List:")

	for i := 1; i < len(records); i++ {
		row := records[i]

		id, _ := strconv.Atoi(row[0])
		desc := row[1]
		created, _ := time.Parse(time.RFC3339, row[2])
		done, _ := strconv.ParseBool(row[3])

		task := Task{
			ID:          id,
			Description: desc,
			CreatedAt:   created,
			IsComplete:  done,
		}

		elapsed := time.Since(task.CreatedAt).Round(time.Minute)
		fmt.Fprintf(w, "%d\t%s\t%s\t%v\n", task.ID, task.Description, elapsed, task.IsComplete)
	}
	w.Flush()

	return nil
}

func CompleteTask(targetID int) error {
	file, err := os.OpenFile("tasks.csv", os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("Error reading CSV: %w", err)
	}

	found := false
	for i := 1; i < len(records); i++ {
		if len(records[i]) < 4 {
			continue // Skip malformed records
		}

		id, _ := strconv.Atoi(records[i][0])
		if id == targetID {
			fmt.Println("Marking task complete:", records[i])
			records[i][3] = "true"
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Task with ID %d not found", targetID)
	}

	// Write the updated records back to the file
	file.Seek(0, 0)
	err = file.Truncate(0)
	if err != nil {
		return fmt.Errorf("Failed to truncate file: %w", err)
	}

	writer := csv.NewWriter(file)
	writer.WriteAll(records)
	writer.Flush() // flush the buffer to disk
	if err := writer.Error(); err != nil {
		return fmt.Errorf("Failed to flush writer: %w", err)
	}

	writer.Flush()
	return nil
}
