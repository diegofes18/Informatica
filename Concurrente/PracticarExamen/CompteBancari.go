package main

import (
	"fmt"       // o log para printear
	"math/rand" // Para generación de numeros aleatorios
	"time"      // sleep
)

// Constantes
const (
	Clientes     = 7 //Clientes que piden al servidor
	Repeticiones = 2 //Numero de veces de la accion

	// Control de velocidades
	cliente_max_time = 5
	cliente_min_time = 1

	//Posibles tipos de las peticiones:
	Positivo = 1
	Negativo = 2
)

// Estructuras
type Empty struct{}

type Deposito struct {
	id       int //Id del proceso
	tipo     int //Estado de peticion
	cantidad int

	// En algunos casos puede ser interesante pasarle un canal en la petición
}

// Variables
var (
	peticionDeposito = make(chan Deposito) //Canal de peticiones al servidor que hacen los clientes
	okDeposito       [Clientes]chan Empty

	peticionConsulta = make(chan int)   //Canal de peticiones al servidor que hacen los clientes
	okConsulta       [Clientes]chan int // Respuestas del servidor a los clientes individualmente

	cliente_nombres = [Clientes]string{"cli1", "cli2", "cli3", "cli4", "cli5", "cli6", "cli7"} // En caso de que los clientes tengan nombres, se accede a este array
)

// Servidor
func servidor() {

	capital := 10000

	for {

		select {

		case peticionDep := <-peticionDeposito:

			switch peticionDep.tipo {

			case Positivo:
				capital = capital + peticionDep.cantidad
				okDeposito[peticionDep.id] <- Empty{}

			case Negativo:
				if capital-peticionDep.cantidad > 0 {
					capital = capital - peticionDep.cantidad
				} else {
					fmt.Println("abortamos")
				}

				okDeposito[peticionDep.id] <- Empty{}
			}

		case id := <-peticionConsulta:
			okConsulta[id] <- capital //

		}

	}
}

func cliente(id int, done chan int) {
	//Inicilizacion del canal de espera personal de cada cliente
	okDeposito[id] = make(chan Empty)
	okConsulta[id] = make(chan int)

	//Variables
	velocidad := cliente_min_time + rand.Intn(cliente_max_time-cliente_min_time)

	//Funciones
	accion1 := func() {
		fmt.Printf("****** %s: Pide Accion 1\n", cliente_nombres[id])

		//Mandamos peticion
		peticionDeposito <- Deposito{id, Positivo, rand.Intn(5000)}
		//Esperamos a que nos confirme
		<-okDeposito[id]
	}

	accion2 := func() {
		fmt.Printf("****** %s: Pide Accion 2\n", cliente_nombres[id])

		//Mandamos peticion
		peticionConsulta <- id
		//Esperamos a que nos confirme
		consultado := <-okConsulta[id]
		fmt.Printf(" Lo consultado por %s, es %s: \n", cliente_nombres[id], consultado)
	}

	//Metodo para esperar
	wait := func(wait_time int) {
		time.Sleep(time.Duration(wait_time) * time.Second)
	}

	//Acciones que hace el Cliente
	for i := 0; i < Repeticiones; i++ {
		accion1()
		wait(velocidad)

		accion2()
		wait(velocidad)

	}
	done <- id
}

func main() {
	rand.Seed(time.Now().UnixNano())

	done := make(chan int) //Sincrono
	//done := make(chan Vacio, 1) //Asincrono

	for i := 0; i < Clientes; i++ {
		go cliente(i, done)
	}
	go servidor()

	for i := 0; i < Clientes; i++ {
		nombre := cliente_nombres[<-done]
		fmt.Printf("*** %s:  Se va\n", nombre)
	}

	fmt.Printf("End\n")
}
