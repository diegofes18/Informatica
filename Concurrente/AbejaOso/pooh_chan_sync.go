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

func producer(i int, jar chan int, mutex chan Empty, done chan Empty) {
	fmt.Printf("Start producer\n")

	for j := 0; j < ToProduce / BEES; j++ {
		time.Sleep(100 * time.Millisecond)

		n := <-jar
		fmt.Printf("%d put honey (jar %d of %d)\n", i, (n + 1), N)

		if n < N - 1 {		// Jar NOT full after put the honey
			n = n + 1
			jar <- n
		} else {			// Jar full after put the honey
			fmt.Printf("WAKE UP POOH!\n")
			mutex <- Empty{}
		}
	}

	fmt.Printf(">>> Bee %d is done\n", i)
	done <- Empty{}
}

func consumer(jar chan int, mutex chan Empty, done chan Empty) {
	fmt.Printf("Pooh is ready to consume %d jars\n", ToProduce / N)

	for i := 0; i < ToProduce / N; i++ {
		<-mutex

		fmt.Printf("\tPooh eat\n")
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("\tRelease jar\n")

		jar <- 0
	}

	fmt.Printf("\tPooh is done\n")
	done <- Empty{}
}

// Initialize jar and provide jar access
func jarproc(ch chan int) {
	ch <- 0
	for {
		n := <-ch
		ch <- n
	}
}

func main() {
	runtime.GOMAXPROCS(Procs)

	jar := make(chan int)
	mutex := make(chan Empty)
	done := make(chan Empty, 1)

	go jarproc(jar)
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
