// package main

// import "fmt"

// // cuando el operado <- va a la derecha de chan, significa que unicamente es de escritura el channel
// func Generator(c chan<- int) {
// 	for i := 1; i <= 10; i++ {
// 		c <- i
// 	}
// 	close(c) //con esto Cerramos un channel
// }

// // cuando el operador <- va a la izquierda de chan significa que solo es de lectura
// func Double(inp <-chan int, out chan<- int) {
// 	for value := range inp {
// 		out <- 2 * value
// 	}
// 	close(out)
// }

// func Print(c <-chan int) {
// 	for value := range c {
// 		fmt.Println(value)
// 	}
// }

// func main() {
// 	generator := make(chan int)
// 	double := make(chan int)

// 	go Generator(generator)
// 	go Double(generator, double)
// 	Print(double)
// }
