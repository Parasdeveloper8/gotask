package main

import (
	"fmt"
	dayone "main/Day1Task/taskfile"
)

func main() {
	// Define the map with Student struct as the key
	student := make(map[dayone.Student]int)
	dayone.AddStudent(student, "paras", 1)
	dayone.AddStudent(student, "shalu", 2)
	dayone.DeleteStudent(student, "paras")
	dayone.DeleteStudent(student, "shalu")
	dayone.RetrieveStudent(student, "paras")
	for index, value := range student {
		fmt.Printf("key %v and value is %v\n", index, value)
	}
}
