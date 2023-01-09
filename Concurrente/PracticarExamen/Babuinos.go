package main

import (
	"fmt"       // o log para printear
	"math/rand" // Para generación de numeros aleatorios
	"time"      // sleep
)

// Constantes
const (
	Babuinos     = 5 //Clientes que piden al servidor
	Repeticiones = 3 //Numero de veces de la accion

	// Control de velocidades
	cliente_max_time = 5
	cliente_min_time = 1
)

// Estructuras
type Empty struct{}

// Variables
var (
	demanaEntrar = make(chan int)
	okEntrar     [Babuinos]chan Empty

	demanaSortir = make(chan int)
	okSortir     = make(chan Empty)

	cliente_nombres = [Babuinos]string{"NORT1", "NORT2", "NORT3", "NORT4", "NORT5"}
)

// Servidor
func servidor() {
	capacidad := 3
	esperan := [Babuinos]int{0, 0, 0, 0, 0}
	numEsperan := 0

	for {

		select {
		case id := <-demanaEntrar:
			if capacidad > 0 {
				capacidad--
				fmt.Println("******* Entra en la cuerda: " + cliente_nombres[id])
				okEntrar[id] <- Empty{}
			} else {
				fmt.Println("******* Espera para entrar: " + cliente_nombres[id] + ", la cuerda esta ocupada")
				esperan[id] = 1
				numEsperan++
			}

		case id := <-demanaSortir:
			//Silla libre:
			capacidad++
			fmt.Println("******* Sale de la cuerda: " + cliente_nombres[id])
			okSortir <- Empty{}

			//Adjudicacion silla a quien espera:
			if numEsperan > 0 {
				//El primero que haya pedido le da una silla
				i := 0
				for i = 0; i < Babuinos; i++ {
					if esperan[i] == 1 {
						break
					}
				}
				capacidad--
				fmt.Println("******* Entra en la cuerda " + cliente_nombres[i] + " ,porque ha salido:  " + cliente_nombres[id])

				okEntrar[i] <- Empty{} //Permiso al que espera
				esperan[i] = 0         //Resetear el valor
				numEsperan--
			}

		}
	}
}

func cliente(id int, done chan Empty) {
	fmt.Println("Hola el meu nom és: " + cliente_nombres[id])

	for i := 1; i <= Repeticiones; i++ {
		//Trabaja
		fmt.Println(cliente_nombres[id] + " :estoy en el norte")
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

		//Pide sentarse
		fmt.Println(cliente_nombres[id] + " quiero cruzar la cuerda")
		demanaEntrar <- id
		<-okEntrar[id] //Espera SU confirmacion

		time.Sleep(time.Duration(rand.Intn(20000)) * time.Millisecond)

		//Pide irse
		fmt.Println(cliente_nombres[id] + " quiero salir de la cuerda")
		demanaSortir <- id
		<-okSortir
	}

	fmt.Println(cliente_nombres[id] + " ya he dado tres vueltas, me duermo")
	done <- Empty{}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	done := make(chan Empty) //Sincrono
	//done := make(chan Vacio, 1) //Asincrono
	for i := range okEntrar {
		okEntrar[i] = make(chan Empty)
	}

	go servidor()

	for i := 0; i < Babuinos; i++ {
		go cliente(i, done)
	}
	for i := 0; i < Babuinos; i++ {
		<-done
	}

	fmt.Printf("End\n")
}
