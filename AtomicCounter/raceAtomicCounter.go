package main

import (
	"runtime"
	"sync/atomic"
)

// setup
const (
	N    = 1000
	Conc = 1000
)

type atomicCounter struct {
	acc int64 // member shared between multiple goroutines, protected by mutex
}

func (c *atomicCounter) addCounter() {
	atomic.AddInt64(&c.acc, 1)
	runtime.Gosched()
}

func (c *atomicCounter) getCounter() int64 {
	return atomic.LoadInt64(&c.acc)
}

func adder(acc *atomicCounter, done chan bool) {
	for i := 0; i < N; i++ {
		acc.addCounter()

	}

	done <- true
}

func main() {
	done := make(chan bool)
	account := atomicCounter{}

	for i := 0; i < Conc; i++ {
		go adder(&account, done)
	}

	// Wait for all goroutines to complete
	for i := 0; i < Conc; i++ {
		<-done
	}

	println(account.getCounter()) // Should == Conc x N
}
