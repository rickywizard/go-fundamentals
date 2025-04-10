package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("title or content cannot be empty")
	}

	return Todo{
		Text: content,
	}, nil
}

func (n Todo) Display() {
	fmt.Printf("\nTodo: %v\n\n", n.Text)
}

func (n Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(n)

	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
