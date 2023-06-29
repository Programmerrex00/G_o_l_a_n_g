package main

import "fmt"

type Employee struct {
	id     int
	nombre string
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.nombre = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.nombre
}

func main() {
	e := Employee{}
	e.SetId(100)
	e.SetName("Maicol")
	//fmt.Printf("%v", e)
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
}
