package week1

import "fmt"

func SliceSum(slice ...int) {
	sum := 0
	for _, num := range slice {
		sum += num
	}
	fmt.Println("Sum: ", sum)
}
