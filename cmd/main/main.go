package main

import (
	"fmt"
	"os"
	"strconv"

	c "github.com/rafaeldiazmiles/GO/pkg/calculator"
)

func main() {

	argumentos := os.Args[1:]
	x, err := strconv.ParseFloat(argumentos[0], 64)
	y, err := strconv.ParseFloat(argumentos[2], 64)

	switch argumentos[1] {
	case "+":
		fmt.Printf("Resultado: %v\n", c.Add(x, y))

	case "-":
		fmt.Printf("Resultado: %v\n", c.Substract(x, y))

	case "x":
		fmt.Printf("Resultado: %v\n", c.Multiply(x, y))

	case "/":
		fmt.Printf("Resultado: %v\n", c.Divide(x, y))

	default:
		fmt.Println(err)
	}
}

//########################################################################
// CODIGO COMENTADO, DE USO PERSONAL (pruebas)
// github.com/rafaeldiazmiles/GO/pkg/project"
// fmt.Println("Chequeo que este activo el main.\n")
// w.Start()
// foo()
// fmt.Println(w.Split(17))
// w.Sum(1, 2)
// nums := []int{1, 2, 3, 4}
// w.Sum(nums...)
// w.SliceSum(nums...)
// w.Index()
// kvs := map[string]string{"a": "apple", "b": "bananas"}
// w.MapIterate(kvs)
// w.OS()
// w.Defer(nums...)
// fmt.Println(w.Sqrt(9))
// w.CastFloat(1)
// w.PointerDemo()
// w.StructDemo()
// w.PSDemo()
// w.PlayerDemo()
// w.ArrayMadness()
// w.MapStructDemo()
// w.FuncTypesDemo()
// w.MethodsDemo()
// w.InterfaceDemo()
// w.TypeAsDemo()
// p.While()
// kvs := map[string]int{"a": 1, "b": 2}

// p.KeySearchStringInt(kvs, "a")
