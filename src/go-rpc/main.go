package main

import (
	"log"
)

// Make a new ToDo type that is a typed collection of fields
// (Title and Status), both of which are of type string
type ToDo struct {
	Title, Status string
}

// Declare variable 'todoSlice' that is a slice made up of
// type ToDo items
var todoSlice []ToDo

// GetToDo takes a string type and returns a ToDo
func GetToDo(title string) ToDo {
	var found ToDo
	// Range statement that iterates over todoArray
	// 'v' is the value of the current iterateee
	for _, v := range todoSlice {
		if v.Title == title {
			found = v
		}
	}
	// found will either be the found ToDo or a zerod ToDo
	return found
}

// MakeToDo takes a ToDo type and appends to the todoArray
func MakeToDo(todo ToDo) ToDo {
	todoSlice = append(todoSlice, todo)
	return todo
}

// EditToDo takes a string type and a ToDo type and edits an item in the todoArray
func EditToDo(title string, editToDo ToDo) ToDo {
	var edited ToDo
	// 'i' is the index in the array and 'v' the value
	for i, v := range todoSlice {
		if v.Title == title {
			todoSlice[i] = editToDo
			edited = editToDo
		}
	}
	// edited will be the edited ToDo or a zeroed ToDo
	return edited
}

// DeleteToDo takes a ToDo type and deletes it from todoArray
func DeleteToDo(todo ToDo) ToDo {
	var deleted ToDo
	for i, v := range todoSlice {
		if v.Title == todo.Title && v.Status == todo.Status {
			// Delete ToDo by appending the items before it and those
			// after to the todoArray variable
			todoSlice = append(todoSlice[:i], todoSlice[i+1:]...)
			deleted = todo
			break
		}
	}
	return deleted
}

func main() {
	log.Println("1. todo Slice: ", todoSlice)
	finishApp := ToDo{"Finish App", "Started"}
	makeDinner := ToDo{"Make Dinner", "Not Started"}
	walkDog := ToDo{"Walk the dog", "Not Started"}
	MakeToDo(finishApp)
	MakeToDo(makeDinner)
	MakeToDo(walkDog)
	log.Println("2. todo Slice: ", todoSlice)
	DeleteToDo(makeDinner)
	log.Println("3. todo Slice: ", todoSlice)
	MakeToDo(makeDinner)
	log.Println("4. todo Slice: ", todoSlice)
	log.Println("5.", GetToDo("Finish App"))
	log.Println("6.", GetToDo("Finish Application"))
	EditToDo("Finish App", ToDo{"Finish App", "Completed"})
	log.Println("7. todo Slice: ", todoSlice)
}