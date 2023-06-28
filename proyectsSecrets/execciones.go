package main

import (
	"fmt"
)

func main() {
	/*myValue, err := strconv.ParseInt("7", 0, 64)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}*/

	/*m := make(map[string]int)
	m["gato"] = 2
	fmt.Println(m["gato"])

	s := []int{1, 2, 3, 4}
	s = append(s, 14, 12)

	for indice, value := range s {
		fmt.Println(indice, value)
	}*/

	/*c := make(chan int)
	go doSomething(c)
	<-c*/

	/*g := 20
	fmt.Println(g)
	h := &g
	fmt.Println(h)*/

	/*c := make(chan int) //Creamos un canal para monitorear las Goroutines
	go doSomething(c)   //Llamamos a la función con go para generar un Goroutine
	<-c                 //main esperara a que este canal reciba el mensaje*/

	g := 25
	fmt.Println(g)
	h := &g         //Apuntador a g
	fmt.Println(h)  //Dirección de memoria donde esta almacenada g
	fmt.Println(*h) //Accedemos al valor de h que a su vez almacena g

}

/*// funcion que simula una gorutine
func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Terminado!!")
	c <- 1
}*/

/*func doSomething(c chan int) { //Recibimos un canal de tipo int
	time.Sleep(3 * time.Second) //Poner a dormir
	fmt.Println("Done")
	c <- 1 // le envia el valor de 1
}*/
