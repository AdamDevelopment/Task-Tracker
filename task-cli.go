
//To-do:
// 1. Add Json content handling - check if record exists, if not, create a new record under first one (CheckIfJsonRecordExists)

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

func ExistsOrCreate(filename string) error {
	if !strings.HasSuffix(filename, ".json") {
		return fmt.Errorf("invalid file extension. Please provide a JSON file")
	}

	// Sprawdzenie, czy plik istnieje
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// Plik nie istnieje, tworzymy nowy plik JSON
			fmt.Println("File does not exist. Creating a new file.")
			err = os.WriteFile(filename, []byte("[]"), 0644) // Pusta tablica JSON jako początkowa zawartość
			if err != nil {
				log.Printf("Error creating file: %v", err)
				return err
			}
			fmt.Println("File created successfully.")
		} else {
			log.Printf("Error checking file existence: %v", err)
			return err
		}
	} else {
		fmt.Println("File exists.")
	}
	return nil
}


// Funkcja do odczytu z pliku JSON
func ReadJsonFile() ([]TaskProperties, error) {
	var tasks []TaskProperties
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return nil, err
	}
	if len(data) == 0 {
		fmt.Println("JSON file is empty.")
		return tasks, nil
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return nil, err
	}
	return tasks, nil
}

func WriteToJsonFile(tasks []TaskProperties, filename string) error {
	data, err := json.MarshalIndent(tasks, "", "   ")
	if err != nil {
		log.Printf("Error marshalling JSON data: %v", err)
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Printf("Error writing to file: %v", err)
		return err
	}
	fmt.Println("Task written to file successfully.")
	return nil
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
	err := ExistsOrCreate(filename)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return
	}

	command := os.Args[1] // add, update, delete, list
	switch command {
	case "add":
		desc := strings.Join(os.Args[2:], " ")
		tasks, err := ReadJsonFile()
		if err != nil {
			log.Printf("Error reading file: %v", err)
			return
		}
		id := len(tasks) + 1 // Generowanie nowego ID
		newTask := AddTask(desc, "Initial status", id)
		tasks = append(tasks, *newTask) // Dodanie nowego zadania do listy

		err = WriteToJsonFile(tasks, filename) // Zapis zaktualizowanej listy zadań do pliku
		if err != nil {
			log.Printf("Error adding a description: %v", err)
			return
		}
	default:
		fmt.Println("Invalid command")
	}

}
