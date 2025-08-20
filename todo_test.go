package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetchTodo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"userId":1,"id":1,"title":"test","completed":false}`)
	}))
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := &http.Client{Timeout: time.Second}
	todo, err := FetchTodo(ctx, client, srv.URL)
	if err != nil {
		t.Fatalf("FetchTodo returned error: %v", err)
	}
	if todo.ID != 1 || todo.Title != "test" || todo.Completed {
		t.Fatalf("unexpected todo: %+v", todo)
	}
}

func TestFetchTodo_DecodeError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "not-json")
	}))
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := &http.Client{Timeout: time.Second}
	if _, err := FetchTodo(ctx, client, srv.URL); err == nil {
		t.Fatal("expected error, got nil")
	}
}
