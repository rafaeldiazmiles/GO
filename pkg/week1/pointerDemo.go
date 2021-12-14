package week1

import "fmt"

func Demo() {
	var someInt *int
	otherInt := 2
	pointer := &otherInt
	fmt.Printf("Zero-value of someInt: %v\n", someInt)
	fmt.Printf("Value pointer points to: %v\n", *pointer)

}
