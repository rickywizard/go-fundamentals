package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// interface in go is going to look after the implemented struct whether it has the method or not
type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputabble interface {
	saver
	Display()
}

// type outputabble interface {
// 	Display()
// 	Save() error
// }

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("hello")

	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	printSomething(todo)

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

	result := add(1, 2)
	fmt.Println(result)
}

// generics
func add[T int | float64 | string](a, b T) T {
	return a + b
}

// any type is interface{}
func printSomething(value any) {
	intVal, ok := value.(int)

	if ok {
		intVal += 1
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		floatVal += 1.5
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		stringVal += "hehe"
		fmt.Println("String:", stringVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// default:
	// 	fmt.Println(value)
	// }
}

func outputData(data outputabble) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Failed saving the todo")
		return err
	}

	fmt.Println("Successfully saved todo")
	return nil
}

func getTodoData() string {
	return getUserInput("Todo text:")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error", err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
