package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
	// can add more function here..
}

func main() {
	title, content := getNoteData()
	text := getTodoData()

	userNote, err := note.New(title, content)
	if err != nil {
		return
	}

	userTodo, err := todo.New(text)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

	err = outputData(userTodo)
	if err != nil {
		return
	}

	printSomething(true)
}

func printSomething(value any) {
	switch value.(type) {
	case int:
		fmt.Println("handling integer data")
	case string:
		fmt.Println("handling string data")
	case bool:
		fmt.Println("handling boolean data")
	}

	fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	return data.Save()
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getTodoData() string {
	text := getUserInput("Todo text:")

	return text
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
