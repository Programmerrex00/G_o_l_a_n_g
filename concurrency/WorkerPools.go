// package main

// import (
// 	"fmt"
// 	"time"
// )

// // LO QUE VAMOS EMULAR EN ESTE PROGRAMA ES SIMULAR COMO ACTUAN LOS WORKERS POOLS QUE SON COMO TRABAJADORES UTILIZANDO CONCURRENCIA
// // UTILIZAREMOS CONCEPTOS COMO LA RECURSIVIDAD, CONCURRENCIA, PARALELISMO, CHANNELS Y CLARAMENTE WORKERS POOLS.
// func main() {
// 	start := time.Now()

// 	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	tas := []int{}
// 	//Ocuparemos solo 8 trabajadores
// 	nWorkers := 8

// 	//definiremos dos channel uno de envio de datos y otro para ver los resultados
// 	jobs := make(chan int, len(tasks))
// 	result := make(chan int, len(tasks))

// 	//Este ciclo for se ejecutara de manera concurrente osea por aparte de la funcion main
// 	// aqui enviaremos el 'id' que seria el indice del ciclo for, por otra parte enviaremos el
// 	// jobs osea el valor resividor del slice tasks, y enviaremos el channel para guardar el resultado
// 	for i := 0; i < nWorkers; i++ {
// 		go Worker(i, jobs, result)
// 	}

// 	//Este for seguira secuencialmete despues que se separa el for de la gorutines y workers, osea el de arriba
// 	// este seguira enviando continuamente valores por el channel 'jobs', simulara una cola
// 	for _, value := range tasks {
// 		jobs <- value
// 	}
// 	//cerramos el channel para no enviar mas valores
// 	close(jobs)

// 	//con este ciclo resiviremos los valores resultante enviados en la funcion Worker por medio del channel results
// 	// osea lo ocupamos para esperar que se completen todas las tareas y resivir el resultado del channel
// 	// ademos podemos guardar los resultados de result en otro slice, para posteriormente mostrarlos
// 	for r := 0; r < len(tasks); r++ {
// 		tas = append(tas, <-result)
// 	}

// 	for i := 0; i < len(tas); i++ {
// 		fmt.Println(tas[i])
// 	}

// 	tiempo := time.Since(start)
// 	fmt.Printf("\nEl programa tardo: %s\n", tiempo)
// }

// // Funcion que se ejecutara concurrentemente Ocupando Workers, osea trabajadores
// func Worker(id int, jobs <-chan int, results chan<- int) {
// 	for job := range jobs {
// 		fmt.Printf("Trabajador con id %d inicia la serie fibonacci con %d\n", id, job)
// 		fib := Fibonacci(job)
// 		fmt.Printf("Trabajador con id %d, trabaja con el numero %d y el resultado es: %d\n", id, job, fib)
// 		results <- fib
// 	}
// }

// func Fibonacci(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	return Fibonacci(n-1) + Fibonacci(n-2)
// }

// //EJEMPLO DE LA EJECUCION CON 8 WORKERS
// // Worker with id 6 started fib with 7
// // Worker with id 6, job 7 and fib 13
// // Worker with id 7 started fib with 8
// // Worker with id 7, job 8 and fib 21
// // Worker with id 0 started fib with 2
// // Worker with id 0, job 2 and fib 1
// // Worker with id 1 started fib with 3
// // Worker with id 1, job 3 and fib 2
// // Worker with id 2 started fib with 1
// // Worker with id 2, job 1 and fib 1
// // Worker with id 3 started fib with 5
// // Worker with id 3, job 5 and fib 5
// // Worker with id 5 started fib with 6
// // Worker with id 5, job 6 and fib 8
// // Worker with id 4 started fib with 4
// // Worker with id 4, job 4 and fib 3
