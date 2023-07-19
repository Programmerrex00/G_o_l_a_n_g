package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

//Programa que ejecuta concurrentemente trabajos entre Workers y Dispatchers, usando
// una serie de canales para la comunicaciòn entre los dos tipos, ademas las tareas que se ejecutaran por los
// workers van ha ser sacar el numero de la serie fibonacci, pero ademas, lo haremos haciendo uso
//de un servidor local, osea usaremos una api rest para enviar solamente solicitudes Por medio del emtodo POST

// struct con atributos como el nombre del trabajo, el tiempo de duración y el numero
// a consultar en la serie fibonacci
type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

// struct que simula la estructura de un trabajador o empleado, en el struct worker definimos
// el id del trabajador, el JobQueue que es el canal donde el resivira los trabajos, el WorKerPool
// que es un canal de canales donde se registran lo worker por medio del canal JobQueue para obtener las tareas
// que hayan sido ya disponibles por el Dispatcher y por ultimo el Quitchan, que es un canal
// que si se recive un valor de true el worker dejara de trabajar
type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

// El dispatcher es el encargado de transmitir esas tareas que resive por medio del metodo POST al los workers
// osea a los trabajadores, tenemos que los atributos son el WorkerPool que es otro canal de canales que es el que se
// encarga de resivir las tareas del canal JobQueue y ponerlas en la lista para que un worker la tome, el MaxWorkers
// que es el encargado de declarar el maximo de trabajadores(workers), para ejecutar las tareas y por ultimo
// JobQueue que es el canal encargado de resivir las tareas y pasarlas al WorkerPool
type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

// En el constructor de Worker, tenemos que resive como parametro un id para el trabajador, y un canal
// de canales, ademas creamos artificalmente dos canales, uno apra JobQueue que es canal donde los workers
// resiven el trabajo y un QuitChan para determinar que el trabajador termino esa tarea. Por ultimo
// retornamos el Worker
func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}
}

// En la funcion Start tenemos varias cosas importantes, primero que es una funcion perteneciente a Worker,
// lo otro es que en el cuerpo de esta funcionn se ejecuta otra funcion anonima y concurrete,
// vemos que se ejecuta un for infinito, despues visulaizamos que el canal en donde resive el trabajao los worker
// lo enviamos al canal de canales osea al WorkerPool, esto lo hacemos para afirmarle al dispatcher, que
// un nuevo canal de worker se ofrese para resivir trabajo, luego que el dispatchar asigne trabajo a ese canal
// validamos con un select si el canal de trabajo del worker ya tiene alguna tarea asignada. Si ese es el caso
// ejecutamos lo que tiene ese case con su respectiva funcion fibonacci. Luego se valida si el QuitChan
// de ese canal tiene el valor de true, si ese es el caso, finaliza la tarea el Worker
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue

			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker with id %d Started\n", w.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d Finished with result %d\n", w.Id, fib)
			case <-w.QuitChan:
				fmt.Printf("Worker with id %d stopped\n", w.Id)
			}
		}
	}()
}

// Funcion para para el trabajo de un worker si en el QuitChan es true
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

// Funcion fibonacci para calcular un numero de esta serie por un valor dato, de manera recursiva
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// El constructor de de Dispatcher, resive como parametro un canal de trabajo y la cantidad de worker osea trabajadores
// que pondra contener. declararemos un worker que lo asignaremos al canal de canales para asignar las caracteristicas
// del worker osea el trabajador, la cantidad maxima 'worker' sera asignada
func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: worker,
	}
}

// El metodo Dispatch se encargara de asignar las tareas obtenida por el canal JobQueue de Dispatcher, osea
// el primero obtendra las tareas y luego las asignara a los worker. Se hara de la siguiente forma:
// Pirmero tendremos una funcion propia de dispatcher, luego en el cuerpo de la funciion tendremos un for que
// se ejecuta infinitamente tendremos un select para multiplexar el trabajo, declararemos una variable 'job'
// para asignarle el trabajo obtenido por JobQueue de dispatcher, luego en una funcion anonima cocurrente
// guardaremos la el canal de canales de dispatcher osea WorkerPool osea se supone que este punto ya hayan
// trabajadores disponibles en el WorkerPool, si ese es el caso entonces se extraera el canal de trabajo de ese worker
// si se guardara en la variable 'workerJobQueue', luego de guardar el canal del worker hay, entonces procederemos
// a asignar el trabajo a ese canal del trabajador worker con 'workerJobQueue <- job'
func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

// En la funcion Run es donde procederemos a hacer el llamado de todas las demas funciones, Primero crearemos
// una funcion propia de Dispatcher llamada run, luego por medio de un for, ya determinano por d.MaxWorkers
// procederemos a inicializar a los worker con el contructor, pasandole como parametro el 'i' id del worker,
// luego, d.WorkerPool, quiere decir que estariamos definiendo el canal de canales de los workers. Luego de eso
// procederemos a iniciar la funcion Start() con los worker ya definidos con sus respectivos ids
// cabe rezaltar que al mismo tiempo que esto va sucediendo, debemos tener en cuenta que mas abajo hay una funcion
// concurrete que tambien se va ejecutando esto quiere decir que al mismo tiempo va enviando tareas al canal de canales
// porque concurrentemente se va ejecutando la funcion Dispatch(), osea si llegan trabajos, al mismo tiempo
// va poniendo a trabajar a los workers(trabajadores)
func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

// La funcion RequestHandler esta hecha para validar diferente cosas la primera es que
// valida que el metodo que sea enviado sea POST, si ese no es el caso se enviara un error
// luego validad cada argumento que conformara al job, osea el tiempo, el valor y el nombre
// luego si esto es verdad crearemos una instancia de job y asignaremos ese trabajo, por otro lado notamos que tenemos
// como paratro un jobQueue, este es un canal y nos ayudara a transportar el job por medio de ese canal
// osea por medio de 'jobQueue' pasamos el job al dispatcher
func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" { //Este mensaje saldra si el usuario envia un GET, PUT, DELETE
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	job := Job{Name: name, Delay: delay, Number: value}
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}

// En la funcion main lo que hacemos es declarar una serie de constantes y incializarlas con el maximos de
//
//	trabajadores 'maxWorkers', la cantidad maxima de trabajos que podra tener un worker 'maxQueueSize'
//
// y el puerto, con esas constantes definidas tenermos que creamos un canal con esas cualidad osea
// del  maxQueueSize que sea de tipo Job, luego crearemos una instancia de dispatcher llamando a su costructor
// y pasando como parametro el jobQueue creado previamente arriba y el maxWokers osea la cantidad de trabajadores
// workers, luego lo siguiente es llamar a la funcion RequestHandler, con su respectivo atributo ya creado
// osea el jobQueue, y log.Fatal(http.ListenAndServe(port, nil)), es para mantener una comucnicacion contstate
// con el servidor, osea para estar escuchandolo
func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8081"
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)
	dispatcher.Run()

	//Buscaremos como http://localhost:8081/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
