package week1

import "fmt"

func Defer(slice ...int) {
	defer fmt.Println("BOOM!")

	for i := range slice {
		defer fmt.Println(i)
	}
	fmt.Println("Counting...: ")
}
