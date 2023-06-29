package main

import "fmt"

type persona struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	persona //campos anonimos
	Employee
}

func GetMessage(p persona) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {
	ejemEmployee := FullTimeEmployee{}
	ejemEmployee.id = 12
	ejemEmployee.age = 23
	ejemEmployee.name = "Maicol"
	fmt.Printf("%v\n", ejemEmployee)
}
