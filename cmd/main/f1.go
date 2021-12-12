package main

import "fmt"

func foo() {
	fmt.Println("Algo enviado directo del main package\n")
}

/*
//copio el material que estuve escribiendo previamente -->

// s := str{}

	// extract(s)
	// extract(2)
	// extract("hello")
	// extract(2.4)
	// s.Greet("gopher")
	// s := str{}
	// s.Greet("gopher")

	// whatType(s)
	// whatType(3)
	// whatType("hi")

	// greetAndBye(s, "gopher")

	// name := mstr("Gopher")
	// name.Hi()
	//
	// otherName := &name
	// otherName.Hi()

	// m := map[string]someStruct{
	// 	"key1": someStruct{2, false},
	// 	"key2": someStruct{3, true},
	// }

	// fmt.Printf("m[\"key2\"].intField=%v\n", m["key2"].intField)

	// m["key3"] = someStruct{4, true}
	// fmt.Printf("m[\"key3\"].boolField=%v\n", m["key3"].boolField)

	// var fun func() error
	// fun = func() error {
	// 	fmt.Println("hi from fun")
	// 	return errors.New("error from fun")
	// }
	// someFunc(fun)
}

func extract(v interface{}) {
	// greeter := v.(str)
	// greeter.Greet("gopher1")
	// gr, ok := v.(str)
	// if !ok {
	// 	fmt.Println("v is not of type str")
	// 	return
	// }
	// gr.Greet("gopher2")
	switch t := v.(type) {
	case str:
		t.Greet("GOPHER")
	case int:
		fmt.Printf("sent an int: %v\n", t)
	case string:
		fmt.Printf("sent a string: %v\n", t)
	default:
		fmt.Printf("unsupported type: %T\n", t)
	}
}

// func someFunc(f func() error) error {
// 	fmt.Println("hi from someFunc")
// 	return f()
// }

// func whatType(v interface{}) {
// 	fmt.Printf("type of %v : %T\n", v, v)
// }

// func greetAndBye(g Greeter, name string) {
// 	g.Greet(name)
// 	fmt.Printf("Bye %v\n", name)
// }




type Greeter interface {
	Greet(name string)
}

type error interface {
	Error() string
}

type str struct{}

func (str) Greet(name string) {
	fmt.Printf("Hi %v\n", name)
}

// type mstr string

// func (m *mstr) Hi() {
// 	fmt.Printf("Hi %v\n", *m)
//}

// type someStruct struct {
// 	intField  int
// 	boolField bool
// }
*/
