// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // Es necesario utilizar WaitGroup para poder utilizar la capacidad definida de las gorutines y hacer parecer un semaforo
// func main() {
// 	c := make(chan int, 2)
// 	var wg sync.WaitGroup
// 	for i := 0; i < 10; i++ {
// 		c <- 1
// 		wg.Add(1)
// 		go doSomething(i, &wg, c)
// 	}
// 	wg.Wait()

// 	// for i := 0; i < 10; i++ {
// 	// 	go func(id int) {
// 	// 		fmt.Printf("Funcion %d Inicio\n", id)
// 	// 		time.Sleep(1 * time.Second)
// 	// 		fmt.Printf("Funcion %d Fin\n", id)
// 	// 	}(i)
// 	// }
// 	// time.Sleep(2 * time.Second)
// }

// func doSomething(i int, wg *sync.WaitGroup, c chan int) {
// 	defer wg.Done()
// 	fmt.Printf("Id: %d started\n", i)
// 	time.Sleep(4 * time.Second)
// 	fmt.Printf("Id: %d finished\n", i)
// 	<-c

// } //Excelente concurrencia
