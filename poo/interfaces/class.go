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
	endDate string
}

type TemporaryEmployee struct {
	persona
	Employee
	taxRate int
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

func (TemEmployee TemporaryEmployee) getMessage() string {
	return "Temporaty Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.id = 12
	ftEmployee.age = 23
	ftEmployee.name = "Maicol"
	fmt.Printf("%v\n", ftEmployee)

	//Uso de las interfaces, simulando el polimorfismo, de manera implicita
	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)

}
