/*package main

import (
	"fmt"
	"sync"
	"time"
)

// Otra manera de ejecutar Concurrentemente un codigo
// Osea lo que hacemos es que con wg.Add(), añadimos 1 al contador WaitGroup y con
// wg.Done() le restamos 1 al contador de WaitGroup, y despues con wg.Wait(), lo que hacemos
// es que espere o bloquee para que todas las gorutines finalizen, cuando lo hagan
// osea cuando el contador sea 0, seguira el proceso regular del programa.
func main() {
	//Declaramos un wg sincrono
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		//Con esto sumanos uno al wg
		wg.Add(1)
		go doSomething(i, &wg)
	}
	//Con esto finalizamos hasta que el contador sea Cero, Espera a que las gorutine Finalizen
	wg.Wait()

	// 		// Canal de tipo int
	// intChan := make(chan int)
	// go func() {
	// 	intChan <- 42 // Enviar un valor al canal
	// }()
	// value := <-intChan // Recibir un valor del canal
	// fmt.Println(value)

	// // Canal de tipo string
	// stringChan := make(chan string)
	// go func() {
	// 	stringChan <- "Hola" // Enviar un valor al canal
	// }()
	// message := <-stringChan // Recibir un valor del canal
	// fmt.Println(message)

	// // Canal de tipo struct
	// type Person struct {
	// 	Name string
	// 	Age  int
	// }
	// personChan := make(chan Person)
	// go func() {
	// 	person := Person{Name: "John", Age: 25}
	// 	personChan <- person // Enviar un valor al canal
	// }()
	// person := <-personChan // Recibir un valor del canal
	// fmt.Println(person)

	// done := make(chan bool)

	// for i := 0; i < 10; i++ {
	// 	go doSomethings(i, done)
	// }

	// // Esperar a que todas las goroutines terminen
	// for i := 0; i < 10; i++ {
	// 	<-done
	// }

	// fmt.Println("Finish")
}

// func doSomethings(i int, done chan<- bool) {
// 	fmt.Printf("Hola soy el mensaje %d Concurrente iniciado\n", i)
// 	time.Sleep(2 * time.Second)
// 	fmt.Printf("Adios el mensaje  %d concurrente ha terminado\n", i)
// 	done <- false
// }

func doSomething(i int, wg *sync.WaitGroup) {
	//Con esto damos por terminada una gorutine y le restamos 1
	defer wg.Done()

	//Imprimimos este mensaje concurrentemente
	fmt.Printf("Started %d\n", i)

	//Esperan todas las gorutines 2 second
	time.Sleep(2 * time.Second)

	//saldra este mensaje concurrentemente en cada gorutine
	fmt.Println("Finish")
}
*/