package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Person
	Employee
}

//(PARA PODER SEPARAR CADA FUNCION LAS DECLARAMOS ASI, COMO FUNCIONES 'ANONIMAS', PARA PODER UTILIZARLAS INDIVIDUALMENTE)
//Simular que se obtiene la persona por medio del DNI, en la base de datos
var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	return Person{}, nil
}

//Simular que se obtiene el empleado por medio del Id, en la base de datos
var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

//Lo que hace esta funcion es verificar si lo que devuelve las dos funciones anteriores es verdad entonces esa persona que es un
// empleado existe, la creacion de estas funciones es para verificar como llevar un test que simula la conexion a base de datos, asi que tenemos que hacer mocks
//para separar esas dos funciones que estan pegadas a GetFullTimeEmployeeById, porque lo que estamos haciendo son pruebas unitarias
func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee = e

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = p

	return ftEmployee, nil
}
