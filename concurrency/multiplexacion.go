// package main

// import (
// 	"fmt"
// 	"time"
// )

// // En este ejecicio vemos que se imprimen dos valores de forma concurrente y respetando el tiempo de cada uno
// // es decir los dos channel se ejecutan en paralelo, y si uno tiene menor tiempo que el otro, entoces procedera a mostrarse primero
// func main() {
// 	c1 := make(chan int)
// 	c2 := make(chan int)
// 	d1 := 4 * time.Second
// 	d2 := 2 * time.Second

// 	go DoSomething(d1, c1, 1)
// 	go DoSomething(d2, c2, 2)

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case channelMsg1 := <-c1:
// 			fmt.Println(channelMsg1)

// 		case channelMsg2 := <-c2:
// 			fmt.Println(channelMsg2)
// 		}
// 	}

// }

// func DoSomething(i time.Duration, c chan<- int, param int) {
// 	time.Sleep(i)
// 	c <- param
// }
