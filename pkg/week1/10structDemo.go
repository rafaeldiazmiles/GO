package week1

import "fmt"

type someStruct struct {
	intField  int
	boolField bool
}

func StructDemo() {
	s := someStruct{}
	l := someStruct{2, true}
	m := someStruct{
		boolField: false,
		intField:  3,
	}

	fmt.Printf("s (valor zero) int field: %v\ns (valor zero) bool field: %v\n", s.intField, s.boolField)
	fmt.Printf("l int field: %v\nl bool field: %v\n", l.intField, l.boolField)
	fmt.Printf("m int field: %v\nm bool field: %v\n", m.intField, m.boolField)

}
