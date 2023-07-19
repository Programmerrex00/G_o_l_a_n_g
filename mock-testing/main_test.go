package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Name: "Maicol Cubides",
						Age:  35,
						DNI:  "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age:  35,
					DNI:  "1",
					Name: "Maicol Cubides",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err != nil {
			t.Errorf("Error when getting Employee")
		}

		if ft.Age != test.expectedEmployee.Person.Age {
			t.Errorf("Error, got %d, expected %d", ft.Age, test.expectedEmployee.Person.Age)
		}
		if ft.DNI != test.expectedEmployee.Person.DNI {
			t.Errorf("Error, got %v, expected %v", ft.Person.DNI, test.expectedEmployee.Person.DNI)
		}
		if ft.Name != test.expectedEmployee.Person.Name {
			t.Errorf("Error, got %v, expected %v", ft.Person.Name, test.expectedEmployee.Person.Name)
		}
		if ft.Id != test.expectedEmployee.Employee.Id {
			t.Errorf("Error, got %v, expected %v", ft.Employee.Id, test.expectedEmployee.Employee.Id)
		}
		if ft.Position != test.expectedEmployee.Employee.Position {
			t.Errorf("Error, got %v, expected %v", ft.Employee.Position, test.expectedEmployee.Employee.Position)
		}
	}
	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
