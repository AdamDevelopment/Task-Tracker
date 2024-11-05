package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"strconv"
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

func WriteToJsonFile(tasks []TaskProperties, filename string) (error, bool) {
	data, err := json.MarshalIndent(tasks, "", "   ")
	if err != nil {
		log.Printf("Error marshalling JSON data: %v", err)
		return err, false
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Printf("Error writing to file: %v", err)
		return err, false
	}
	return err, true
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

func UpdateTask(tasks []TaskProperties, id int, desc string) {
	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Desc = desc
			tasks[i].UpdatedAt = time.Now()
		}
	}
}

func DeleteTask(tasks []TaskProperties, id int) []TaskProperties {
	for i, task := range tasks {
		if task.Id == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}
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

		err, executed  := WriteToJsonFile(tasks, filename) // Zapis zaktualizowanej listy zadań do pliku
		if err != nil || !executed {
			log.Printf("Error adding a description: %v", err)
			return
		} else {
			fmt.Println("Task added successfully.")
		}

	case "update":
		id := os.Args[2]
		idStr, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Error converting ID to integer: %v", err)
			return
		}
		desc := strings.Join(os.Args[3:], " ")
		tasks, err := ReadJsonFile()
		if err != nil {
			log.Printf("Error reading file: %v", err)
			return
		}
		UpdateTask(tasks, idStr, desc)
		err, executed := WriteToJsonFile(tasks, filename)
		if err != nil || !executed {
			log.Printf("Error updating task: %v", err)
			return
		} else {
			fmt.Println("Task updated successfully.")
		}
	case "delete":
		// id := os.Args[2]
		// idStr, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Error converting ID to integer: %v", err)
			return
		}
		tasks, err := ReadJsonFile()
		if err != nil {
			log.Printf("Error reading file: %v", err)
			return
		}
		// DeleteTask(tasks, idStr)
		log.Println(tasks)
		err, executed := WriteToJsonFile(tasks, filename)
		log.Println(tasks)
		if err != nil || !executed {
			log.Printf("Error deleting task: %v", err)
			return
		} else {
			fmt.Println("Task deleted successfully.")
		}
	default:
		fmt.Println("Invalid command")
	}

}
