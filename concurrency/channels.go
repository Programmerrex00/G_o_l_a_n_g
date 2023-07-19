/*package main

import "fmt"

func main() {
	//declaracion del canal que transmite enteros, de tamaño de 3, si no le pongo tamaño se bloqueara el canal
	//el tamaño que el declaramos se llama buffe
	c := make(chan int, 3)

	//Envio informacion por ese channel
	c <- 1
	c <- 1112121
	c <- 1112121432

	//Imprimimos lo que ese channel trae
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	//Asi podemos utilizar un channel unbuffered, y podemos enviar cuantos datos querramos
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 432
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}*/
