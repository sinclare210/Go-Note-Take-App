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
