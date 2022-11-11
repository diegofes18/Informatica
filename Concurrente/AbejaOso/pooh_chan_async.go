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
)

type Empty struct{}

func producer(i int, jar chan int, mutex chan Empty, done chan bool) {
	fmt.Printf("Start producer\n")

	for j := 0; j < ToProduce / BEES; j++ {
		time.Sleep(100 * time.Millisecond)

		n := <-jar
		fmt.Printf("%d put honey (%d de 10)\n", i, (n + 1))

		if n < N - 1 {		// Jar NOT full after put the honey
			n = n + 1
			jar <- n 
		} else {			// Jar full after put the honey
			fmt.Printf("WAKE UP POOH!\n")
			mutex <- Empty{}
		}
	}
	done <- true
}

func consumer(jar chan int, mutex chan Empty, done chan bool) {
	fmt.Printf("Start consumer\n")

	for i := 0; i < ToProduce / N; i++ {
		<-mutex

		fmt.Printf("\tPooh eat\n")
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("\tRelease jar\n")

		jar <- 0
	}

	fmt.Printf("\tPooh is done\n")
	done <- true
}

func main() {
	runtime.GOMAXPROCS(Procs)

	jar := make(chan int, 1)
	jar <- 0
	mutex := make(chan Empty, 1)
	done := make(chan bool, 1)

	go consumer(jar, mutex, done)
	for i := 0; i < BEES; i++ {
		go producer(i, jar, mutex, done)
	}

	for i := 0; i < BEES; i++ {
		<-done
	}
	<-done

	fmt.Printf("DONE\n")
}
