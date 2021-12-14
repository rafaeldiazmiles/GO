package week1

import "fmt"

type User struct {
	ID             int
	Name, Location string
}

type Player struct {
	User
	GameID int
}

func PlayerDemo() {
	p := Player{}
	p.ID = 42
	p.Name = "Globant"
	p.Location = "Montevideo"
	p.GameID = 4949
	fmt.Printf("%+v\n", p)
}
