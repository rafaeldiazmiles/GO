package week1

import "fmt"

func Index() {
	numeros := []int{1, 2, 3, 4}
	for i, num := range numeros {
		if num == 3 {
			fmt.Println("El valor de i: ", i)
		}
	}
}
