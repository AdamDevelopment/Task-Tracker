package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	//"flag"
	"encoding/json"
)

var (
	filename string
)

type TaskProperties struct {
	Id        int       `json:"id"`          // A unique identifier for the task
	Desc      string    `json:"description"` // A description of the task
	CreatedAt time.Time `json:"createdAt"`   // The time the task was created
	UpdatedAt time.Time `json:"updatedAt"`   // The time the task was last updated
	Status    string    `json:"status"`      // The status of the task
}

func JsonDataMarshalling(t *TaskProperties) ([]byte, error) {
	return json.MarshalIndent(t, "", "   ")
}

func WriteJsonFile(t *TaskProperties) error {
	data, err := JsonDataMarshalling(t)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
		return err
	}
	fmt.Print("Enter the file name: ")
	_, err = fmt.Scanln(&filename)

	if !strings.HasSuffix(filename, ".json") {
		filename += ".json"
	} else {
		fmt.Println("File name already has a .json extension")
	}

	if err != nil {
		log.Fatalf("Error reading file name: %v", err)
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
		return err
	}
	return nil
}

// to-do: allow user to input the task properties
func main() {
	t := TaskProperties{
		Id:        1,
		Desc:      "Learn Go",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    "Open",
	}
	err := WriteJsonFile(&t)
	if err != nil {
		log.Fatalf("Error writing JSON to file: %v", err)
	}
	fmt.Printf("Data written to file %s\n", filename)

}
