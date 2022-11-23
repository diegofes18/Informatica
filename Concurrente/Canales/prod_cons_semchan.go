package main

import (
    "fmt"
    "runtime"
)

const (
    Procs      = 4
    MaxCount   = 10000000
    Goroutines = 4
)

type Empty struct{}
type Semaphore chan Empty

func NewSemaphore(value int) Semaphore {
    s := make(Semaphore, 256) // Max
    for i := 0; i < value; i++ {
        s <- Empty{}
    }
    return s
}

func (s Semaphore) Wait() {
    <-s //agafa missatge del canal
}

func (s Semaphore) Signal() {
    s <- Empty{}
}



