
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

func ExistsOrCreate(t *TaskProperties, filename string) error {
	if !strings.HasSuffix(filename, ".json") {
		return fmt.Errorf("invalid file extension. Please provide a JSON file")
	}
	data, err := json.MarshalIndent(t, "", "   ")
	if err != nil {
		log.Printf("Error marshalling JSON data: %v", err)
		return err
	}
	_, err = os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {

			SleepTime := 1.5
			fmt.Println("File does not exist.")
			time.Sleep(time.Duration(SleepTime) * time.Second)
			fmt.Print("Creating a new file")
			for i := 0; i < 3; i++ {
				fmt.Print(".")
				time.Sleep(1 * time.Second)

			}
			fmt.Println("")
			err = os.WriteFile(filename, data, 0644)
			if err != nil {
				log.Printf("Error writing to file: %v", err)
				return err
			}
			fmt.Println("File created successfully.")
		} else {
			log.Printf("Error checking file existence: %v", err)
			return err
		}
	} else {
		fmt.Println("File exists")
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

func CheckIfJsonRecordExists(filename string) bool {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return false
	}
	if len(data) == 0 {
		fmt.Println("JSON file is empty.")
		return false
	}
	fmt.Println("JSON file is not empty.")
	return true
}

func WriteToJsonFile(task *TaskProperties, filename string) error {
	data, err := json.MarshalIndent(task, "", "   ")
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
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./task-cli <command> [arguments]")
		return
	}
	err := ExistsOrCreate(&TaskProperties{}, filename)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return
	}

	command := os.Args[1] // add, update, delete, list
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
		// if CheckIfJsonRecordExists(filename) {

		// }else{

		// }
		err = WriteToJsonFile(newTask, filename) // Zapisz nowe zadanie do pliku
		if err != nil {
			log.Printf("Error adding a description: %v", err)
			return
		}
		fmt.Println("Task added successfully")
		// 	// case "update":
		// // 	taskId := os.Args[2]
		// // 	status := os.Args[3]
		// // 	desc := strings.Join(os.Args[3:]," ")
		// // 	if len(os.Args) < 3 {
		// // 		fmt.Println("Please provide a status")
		// // 		return
		// // 	}

		// // 	if err != nil {
		// // 		log.Printf("Error changing status: %v", err)
		// // 		return err
		// // 	}
		// default:
		// 	fmt.Println("Invalid command")
	}

}
