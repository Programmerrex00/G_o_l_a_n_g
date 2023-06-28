package main

import (
	"fmt"
	"time"
)

func main() {

	//c := make(chan int)
	//go doSomething(c)
	//<-c

	g := 20
	fmt.Println(g)
	h := &g
	fmt.Println(h)
}

// funcion que simula una gorutine
func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Terminado!!")
	c <- 1
}
