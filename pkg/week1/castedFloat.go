package week1

import "fmt"

func CastFloat(x int) float64 {
	castedFloat := float64(x)
	fmt.Printf("Type sent: %T\nCasted to: %T\n", x, castedFloat)
	return castedFloat
}
