package main

import (
    "fmt"
    "runtime"
)

const (
    
    eatCount  = 2
    numAbelles  = 6
    POT = 10
)

type Empty struct{}

func os(id int, done chan Empty, buffer chan string) {
    for i := 0; i < eatCount; i++ {
        data := <-buffer
        fmt.Printf("el Oso: %d     consumeix: %s\n", id, data)
    }
    done <- Empty{}
}

func producer(id int, done chan Empty, buffer chan string) {
    for i := 0; i < ToProduce; i++ {
        data := fmt.Sprintf("%d", i)
        buffer <- data
        fmt.Printf("%d produeix: %s\n", id, data)
    }
    done <- Empty{}
}

func server(id int, done chan Empty, buffer chan string) {
    
}

func main() {
    done := make(chan Empty, 1)
    
    buffer := make(chan string, BufferSize)

    for i := 0; i < Producers; i++ {
        go producer(i, done, buffer)
    }

    for i := 0; i < Consumers; i++ {
        go consumer(i, done, buffer)
    }

    for i := 0; i < Consumers + Producers; i++ {
        <-done
    }

    fmt.Printf("End\n")
}
