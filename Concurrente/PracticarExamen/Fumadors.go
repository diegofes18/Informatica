package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	clients = 14
)

type Empty struct{}

type PeticionEntrar struct {
	id   int
	tipo string
}
type PeticionSalir struct {
	id   int
	sala int
}
type Sala struct {
	capacidad int
	tipo      string
}

// Problema basado en Servidor:
// Para cada accion: Peticion - Confirmacion
// Sentarse
var demanaCadira = make(chan PeticionEntrar)
var permisCadira [clients]chan int //El servidor tiene que saber a quien darle la silla
// La forma de indentificarlo es que cada uno tiene uno
// Irse
var demanaSortir = make(chan PeticionSalir)
var permisSortir = make(chan Empty)

// Clase ENANO
func nan(id int, fuma string, done chan Empty) {
	fumador := fuma

	fmt.Println("Hola el meu nom és: " + id)

	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

	//Pide sentarse
	fmt.Println(id + " ha arribat de la mina i espera una cadira")
	demanaCadira <- PeticionEntrar{id, fumador}
	sala := <-permisCadira[id] //Espera SU confirmacion

	fmt.Println(id + " ja seu pren el menu del dia")

	time.Sleep(time.Duration(rand.Intn(20000)) * time.Millisecond)

	//Pide irse
	fmt.Println(id + " ha acabat de menjar i demana permís per aixecar-se")
	demanaSortir <- PeticionSalir{id, sala}
	<-permisSortir

	fmt.Println(nom + " se'n va ")
	done <- Empty{}
}

// Clase MAJORDOMO
func majordom() {

	salas := [3]Sala{}
	for i := range salas {
		salas[i] = Sala{3, "Lliure"}
	}
	esperen := [Clients]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	numEsperen := 0

	for {
		select {

		case Peticion := <-demanaCadira:
			asigna:=0
			
			for i := salas{
				if(salas[i].capacidad >0 and 
					(salas[i].tipo == "Lliure" or salas[i].tipo == Peticion.tipo)){
					asigna = i
				}
			}
			if(asigna != 0){
				salas[asigna] = Sala{salas[asigna].capacidad-1,Peticion.tipo}
				permisCadira[Peticion.id] <- asigna
			}else{
				esperen[Peticion.id]=1
				numEsperen++
			}
			

		case Peticion := <-demanaSortir:

			//Silla libre:
			if salas[Peticion.sala].capacidad == 1{
				salas[Peticion.sala]=Sala{salas[Peticion.sala].capacidad-1,"Lliure"}
			}else{
				salas[Peticion.sala]=Sala{salas[Peticion.sala].capacidad-1,salas[Peticion.sala].tipo}
			}
			fmt.Println("******* El majordom dona permís per anar-se'n ")
			permisSortir <- Empty{}

			//Adjudicacion silla a quien espera:
			if numEsperen > 0 {
				//El primero que haya pedido le da una silla
				i := 0
				for i = 0; i < clients; i++ {
					if esperen[i] == 1 {
						break
					}
				}
				cadires--
				fmt.Println("******* El majordom fa seure a " + nom[i] + " a la cadira de " + nom[id])

				permisCadira[i] <- Empty{} //Permiso al que espera
				esperen[i] = 0             //Resetear el valor
				numEsperen--
			}
		}

	}
}

func main() {
	done := make(chan Empty, 1) //Canal para finalizar

	for i := range permisCadira {
		permisCadira[i] = make(chan Empty)
	}

	go majordom()
	for i := 0; i < Nans; i++ {
		go nan(nom[i], i, done)
	}
	for i := 0; i < Nans; i++ {
		<-done
	}

	fmt.Printf("Simulació acabada\n")
}
