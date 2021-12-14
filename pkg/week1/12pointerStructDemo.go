package week1

import "fmt"

func PSDemo() {
	p := &someStruct{2, false}
	fmt.Printf("p.intField: %v\n", p.intField)
	fmt.Printf("(*p).boolField: %v\n", (*p).boolField)
}
