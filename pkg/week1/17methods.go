package week1

import "fmt"

type mstr string

func (m *mstr) Hi() {
	fmt.Printf("Hi %v\n", *m)
}

func MethodsDemo() {
	name := mstr("Gopher")
	name.Hi()

	otherName := &name
	otherName.Hi()
}
