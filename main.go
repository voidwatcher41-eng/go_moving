package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// main fetches a Todo from a remote API, updates it and prints both the
// original and updated representations.
func main() {
	// Create a context with timeout to avoid hanging requests.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use an HTTP client with a sensible timeout.
	client := &http.Client{Timeout: 5 * time.Second}

	todo, err := FetchTodo(ctx, client, "https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalf("fetch todo: %v", err)
	}
	fmt.Printf("Fetched Todo: %+v\n", todo)

	// Modify the Todo and marshal it back to JSON.
	todo.Completed = true
	updated, err := json.MarshalIndent(todo, "", "  ")
	if err != nil {
		log.Fatalf("marshal todo: %v", err)
	}
	fmt.Printf("Updated JSON: %s\n", updated)
}
