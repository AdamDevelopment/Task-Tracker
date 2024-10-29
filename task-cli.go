package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	filename = "tasks.json"
)

type TaskProperties struct {
	Id        int       `json:"id"`          // A unique identifier for the task
	Desc      string    `json:"description"` // A description of the task
	CreatedAt time.Time `json:"createdAt"`   // The time the task was created
	UpdatedAt time.Time `json:"updatedAt"`   // The time the task was last updated
	Status    string    `json:"status"`      // The status of the task
}

// Funkcja do serializacji obiektu zadania do JSON
func JsonDataMarshalling(t *TaskProperties) ([]byte, error) {
	return json.MarshalIndent(t, "", "   ")
}
func ExistsOrCreate(t *TaskProperties, filename string) error {
	if !strings.HasSuffix(filename, ".json") {
		return fmt.Errorf("file name must end with .json extension")
	}

	data, err := JsonDataMarshalling(t)
	if err != nil {
		log.Printf("Error marshalling JSON data: %v", err)
		return err
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("File does not exist. Creating a new file...")
		err = os.WriteFile(filename, data, 0644)
		if err != nil {
			log.Printf("Error writing to file: %v", err)
			return err
		}
		fmt.Println("File created successfully.")
	} else {
		fmt.Println("File exists.")
	}
	return nil
}


// Funkcja do odczytu z pliku JSON
func ReadJsonFile() (*TaskProperties, error) {
	var task TaskProperties
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return nil, err
	}
	err = json.Unmarshal(data, &task)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return nil, err
	}
	return &task, nil
}

// Funkcja do dodania nowego zadania
func AddTask(desc string, status string, id int) *TaskProperties {
	return &TaskProperties{
		Id:        id,
		Desc:      desc,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    status,
	}
}

// Correct "Add" function and reading + writing to JSON (not working due to wrong order of arguments)
// add "update" and "delete" functions
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./task-cli <command> [arguments]")
		return
	}
	command := os.Args[1] // add, update, delete, list
	err := ExistsOrCreate(&TaskProperties{}, filename)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return
	}
	switch command {
	case "add":
		desc := strings.Join(os.Args[2:], " ") // ie. "Task description"
		task, err := ReadJsonFile()
		if err != nil {
			log.Printf("Error reading file: %v", err)
			return
		}
		id := task.Id + 1 // ID dla nowego zadania
		newTask := AddTask(desc, "Initial status", id)
		err = ExistsOrCreate(newTask, filename) // Zapisz nowe zadanie do pliku
		if err != nil {
			log.Printf("Error adding a description: %v", err)
			return
		}
		fmt.Println("Task added successfully")
		// case "update":
	// 	taskId := os.Args[2]
	// 	status := os.Args[3]
	// 	desc := strings.Join(os.Args[3:]," ")
	// 	if len(os.Args) < 3 {
	// 		fmt.Println("Please provide a status")
	// 		return
	// 	}

	// 	if err != nil {
	// 		log.Printf("Error changing status: %v", err)
	// 		return err
	// 	}
	default:
		fmt.Println("Invalid command")
	}
}
