package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(promtext string) string {
	fmt.Print(promtext)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func GetNoteData() (string, string) {
	title := GetUserInput("Note Title: ")
	content := GetUserInput("Note Content: ")
	return title, content

}

func GetTodoInput(promptext string) string {
	fmt.Print(promptext)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text

}

func GetTodoData() string {
	content := GetTodoInput("Input The Todo: ")
	return content
}
