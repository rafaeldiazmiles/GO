package main

import (
	"fmt"

	w "github.com/rafaeldiazmiles/GO/pkg/week1" //Cuando estoy trabajando con un repositorio en github puedo apuntar a mi ruta de paquetes con la url del repositorio, la w que pongo antes de comenzar es el nombre que le asigno al paquete para llamarlo desde main
)

func main() {
	fmt.Println("Onward i'll be calling every package with a single function, just to check they are working on main\n")
	w.Start()
	foo()
	// fmt.Println(w.Split(17))
}
