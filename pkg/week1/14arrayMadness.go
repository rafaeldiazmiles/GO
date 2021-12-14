package week1

import "fmt"

func ArrayMadness() {
	var someArray [3]int                                     //Array's correct way of declaration
	var names []string                                       //nil value (no array alocated)
	otherNames := make([]string, 2)                          //Initialized with a size of 2 (nils?) and unlimited capacity
	capacity := make([]string, 2, 4)                         //Initialized with a size of 2 (nils?) and 4 maximum capacity
	numbers := []int{1, 2, 3, 4, 5}                          //Slice (or Array?) declared with values
	oneToThree := numbers[0:2]                               //A slice of numbers Array using indexes
	threeToFive := numbers[2:]                               //If we omit one it will go to the last (omit right) or the first (omit left)
	oneToThree = append(oneToThree, 4)                       //With append I can increment the size
	fiveToSix := append(threeToFive[len(threeToFive)-1:], 6) //Aaand: Madness...
	fmt.Println(someArray)
	fmt.Println(names)
	fmt.Println(otherNames)
	fmt.Println(capacity)
	fmt.Println(oneToThree)
	fmt.Println(threeToFive)
	fmt.Println(fiveToSix)
}
