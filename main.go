package main

import (
	"fmt"
	"github.com/sinclare210/Note-Taking-App/Note"
	"github.com/sinclare210/Note-Taking-App/Todo"
)

func main() {
	title, content := GetNoteData()

	text := GetTodoData()
	userTodo, err := todo.New(text)

	userTodo.Display()
	userTodo.Save()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println()
		fmt.Println("Svaing the note failed")
		return
	}
	fmt.Println()
	fmt.Println("Saving the note suceeded")
}

// type User struct {
// 	name string
// 	age int
// 	university string
// }

// func New (name string, age int, university string)(*User){
// 	return &User{
// 		name: name,
// 		age: age,
// 		university: university,
// 	}
// }

// func (user *User) displayName(){
// 	fmt.Println("I am %v from %v, and i am %v", user.name, user.university, user.age)
// }
