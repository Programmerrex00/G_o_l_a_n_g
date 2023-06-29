package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

//Contructor de objetos en Golang, haciendo uso de punteros, devolviendo una referencia de memoria del objeto creado
func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	//1
	e1 := Employee{}
	fmt.Printf("%v\n", e1)

	//2
	e2 := Employee{
		id:       100,
		name:     "Maicol",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)

	//3 utilizamos el * para referirnos al valor como tal, osea haciendo uso de apuntadores
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 111
	e3.name = "bing"
	e3.vacation = true
	fmt.Printf("%v\n", *e3)

	//4
	e4 := NewEmployee(1, "Rom", true)
	fmt.Printf("%v\n", *e4)

}
