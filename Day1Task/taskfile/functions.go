package dayone

import (
	"fmt"
)

type Student struct {
	Name string
}

func AddStudent(m map[Student]int, name string, grade int) {
	m[Student{name}] = grade
	fmt.Println("Student added")
}
func RetrieveStudent(r map[Student]int, key string) {
	data := Student{key}
	_, found := r[data]
	if found {
		fmt.Println(data)
	} else {
		fmt.Println("No student found")
	}
}
func DeleteStudent(d map[Student]int, keytodel string) {
	delete(d, Student{keytodel})
	fmt.Printf("Key deleted %v\n", keytodel)
}
