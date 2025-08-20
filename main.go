package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Todo represents a task from the JSONPlaceholder API.
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// Perform GET request to fetch a Todo item.
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalf("error fetching todo: %v", err)
	}
	defer resp.Body.Close()

	// Read response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response: %v", err)
	}

	// Unmarshal JSON into Todo struct.
	var todo Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		log.Fatalf("error unmarshaling JSON: %v", err)
	}
	fmt.Printf("Fetched Todo: %+v\n", todo)

	// Modify the Todo and marshal back to JSON.
	todo.Completed = true
	updated, err := json.Marshal(todo)
	if err != nil {
		log.Fatalf("error marshaling JSON: %v", err)
	}
	fmt.Printf("Updated JSON: %s\n", updated)
}
