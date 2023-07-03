// package main

// import (
// 	"fmt"
// )

// // Funciones anonimas las uso cuando no quiero crear una funcion afuera
// // para poder llamar a la funcion anonima, podremos () en la llave de cierre

// // La forma de declarar una funcion anonima y volverla a llamar otra vez si crear otra es
// // creandola pero sin los parentesis en la ultima llave
// func main() {
// 	x := 5
// 	y := func(x int) int {
// 		return x * 2
// 	}
// 	z := 40
// 	fmt.Println(y(x))
// 	fmt.Println(y(z))
// 	//Uso de una funcion anonima con gorutines
// 	// c := make(chan int)

// 	// go func() {
// 	// 	fmt.Println("Starting Function")
// 	// 	time.Sleep(5 * time.Second)
// 	// 	fmt.Println("End")
// 	// 	c <- 1
// 	// }()
// 	// <-c
// }
