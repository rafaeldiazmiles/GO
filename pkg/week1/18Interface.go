package week1

import "fmt"

type Greeter interface {
	Greet(name string)
}

type str struct{}

func (str) Greet(name string) {
	fmt.Printf("Hi %v\n", name)
}

func whatType(v interface{}) {
	fmt.Printf("Type of %v: %T\n", v, v)
}

func greetAndBye(g Greeter, name string) {
	g.Greet(name)
	fmt.Printf("Bye %v\n", name)
}

func InterfaceDemo() {
	s := str{}
	s.Greet("Gopher")
	whatType(s)
	whatType(3)
	whatType("Hi")
	greetAndBye(s, "Gopher")
}
