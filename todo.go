package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Todo represents a task returned by the JSONPlaceholder API.
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// FetchTodo retrieves a Todo from the provided URL using the supplied HTTP
// client and context.
func FetchTodo(ctx context.Context, client *http.Client, url string) (Todo, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Todo{}, fmt.Errorf("create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return Todo{}, fmt.Errorf("perform request: %w", err)
	}
	defer resp.Body.Close()

	var todo Todo
	if err := json.NewDecoder(resp.Body).Decode(&todo); err != nil {
		return Todo{}, fmt.Errorf("decode response: %w", err)
	}
	return todo, nil
}
