package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	Procs = 4
	ToProduce = 40
	N = 10
	BEES = 4

	// Requests status for the bees
	ProducingHoney = iota	// Bee is producing honey
							// dont want to interact with the jar
	WantPutHoney			// Bee want to put honey into the jar
							// sent request and wait for permision

	// Requests status for Pooh
	Hungry					// Pooh is hungry and want eat from the jar
	DoneEating				// Pooh is done with the jar and release it
	Dummy					// Dummy status to print last "Pooh release jar" message
)

type Empty struct{}

type Request struct {
	c		chan Empty
	id		int
	status	int
}

func provider(channelBees chan Request, channelPooh chan Request) {
	var n = 0
	var maya_friends [BEES]Request	// Save a list of bees
	var pooh Request				// Save a pooh request

	// Initialize the bees status with "ProducingHoney" and pooh with "DoneEating"
	for i := range maya_friends {
		maya_friends[i].status = ProducingHoney
	}
	pooh.status = DoneEating

	for {
		select {
			
			// Process messages from the bees
		case message_bees := <-channelBees:
			// Save bee request
			maya_friends[message_bees.id] = message_bees
			switch message_bees.status {
			case WantPutHoney:		// Bee want put honey into the jar

				if n < N {			// Jar not full
					// Increment honey count
					n = n + 1
					fmt.Printf("%d put honey (%d de 10)\n", message_bees.id, n)
					// Provide permission to the bee
					message_bees.c <- Empty{}

					// If jar is full (after the increment) and
					// we have a pendent pooh request to eat, give permission
					if n == N && pooh.status == Hungry {
						pooh.c <- Empty{}
					}
				} else {			// Jar is full
					fmt.Printf("Bee %d blocked\n", message_bees.id)
				}
			}

			// Process messages from Pooh
		case message_pooh := <-channelPooh:
			// Save pooh request
			pooh = message_pooh
			switch message_pooh.status {
			case Hungry:			// Pooh want eat honey
				// If the jar is full when pooh request
				// give the permission to eat,
				// otherwise the request is saved and will be processed
				// when the jar is full
				if n == N {
					pooh.c <- Empty{}
				}
			case DoneEating:		// Pooh is done with the honey
				fmt.Printf("\tPooh release jar\n")
				n = 0				// Empty jar
				// For each bee that "want put honey into the jar"
				// give the permission
				for i := range maya_friends {
					if maya_friends[i].status == WantPutHoney {
						n = n + 1
						fmt.Printf(">>> %d put honey (%d de 10)\n", i, n)
						maya_friends[i].c <- Empty{}
					}
				}
			}
		}
	}
}

func producer(i int, channel chan Request, done chan Empty) {
	prodChan := make(chan Empty)
	fmt.Printf("Bee %d is ready to produce %d units of honey\n", i, ToProduce / BEES)

	for j := 0; j < ToProduce / BEES; j++ {
		// Produce honey
		time.Sleep(100 * time.Millisecond)
		// Put honey
		channel <- Request{id: i, c: prodChan, status: WantPutHoney}
		<-prodChan
		// Set Producing
		channel <- Request{id: i, c: prodChan, status: ProducingHoney}
	}

	done <- Empty{}
}

func consumer(channel chan Request, done chan Empty) {
	myChan := make(chan Empty)
	fmt.Printf("Pooh is ready to consume %d jars\n", ToProduce / N)

	for i := 0; i < ToProduce / N; i++ {
		// Want honey
		channel <- Request{id: i, c: myChan, status: Hungry}
		<-myChan
		// Eating
		fmt.Printf("\tPooh eating jar %d of %d\n", i + 1, ToProduce / N)
		time.Sleep(100 * time.Millisecond)
		// Done eating
		channel <- Request{id: i, c: myChan, status: DoneEating}
		// Dummy request to print DoneEating message "Pooh release jar"
		channel <- Request{id: i, c: myChan, status: Dummy}
	}

	done <- Empty{}
}

func main() {
	runtime.GOMAXPROCS(Procs)
	done := make(chan Empty, 1)
	PCBees := make(chan Request)
	PCPooh := make(chan Request)

	// Start provider, consumer and producers
	go provider(PCBees, PCPooh)
	go consumer(PCPooh, done)
	for i := 0; i < BEES; i++ {
		go producer(i, PCBees, done)
	}

	// Wait for bees and pooh
	for i := 0; i < BEES; i++ {
		<-done
	}
	<-done

	fmt.Printf("DONE\n")
}
