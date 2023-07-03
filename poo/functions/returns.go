package main

import (
	"fmt"
)

//crear funciones con parametros dinamicos, y retornos con nombres

func sum(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func printNames(name ...string) {
	for _, v := range name {
		fmt.Println(v)
	}
}

//funcion con retorno de NULL

func sumar(values ...int) (int, error) {
	if len(values) > 0 {
		total := 0
		for _, v := range values {
			total += v
		}
		return total, nil
	} else {
		return 0, fmt.Errorf("Funcion sin argumentos")
	}
}

// retornos con nombres
func getValues(x int) (double int, triple int, quadruple int) {
	double = 2 * x
	triple = 3 * x
	quadruple = 4 * x
	return
}

func main() {
	fmt.Println(sum(5, 5, 4, 4, 4, 4, 4))
	printNames("rexmax", "adolfo", "Maicol")
	fmt.Println(getValues(1))
	fmt.Println(sumar())
}
