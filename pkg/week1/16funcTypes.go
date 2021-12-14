package week1

import (
	"errors"
	"fmt"
)

func FuncTypesDemo() {
	var fun func() error
	fun = func() error {
		return errors.New("error from fun")
	}
	someFunc(fun)
}

func someFunc(f func() error) error {
	fmt.Println(f())
	return f()
}
