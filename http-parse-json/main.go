package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/2")

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		fmt.Printf("Invalid HTTP Status Code: %d\nBody: %s", response.StatusCode, body)
		os.Exit(1)
	}

	var todo Todo

	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("UserID: %d, ID: %d, Title: %s, Completed: %v", todo.UserId, todo.Id, todo.Title, todo.Completed)
}
