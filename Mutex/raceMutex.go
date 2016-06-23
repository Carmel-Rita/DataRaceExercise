package main

import (
	"sync"
)

// setup
const (
	N    = 1000
	Conc = 1000
)

type shared struct {
	mutex *sync.Mutex
	acc   int64 // member shared between multiple goroutines, protected by mutex
}

func (s *shared) addCounter() {
	defer s.mutex.Unlock()
	s.mutex.Lock()
	s.acc += 1
}

func (s *shared) getCounter() int64 {
	defer s.mutex.Unlock()
	s.mutex.Lock()
	return s.acc
}

func adder(acc *shared, done chan bool) {
	for i := 0; i < N; i++ {
		acc.addCounter()

	}

	done <- true
}

func main() {
	done := make(chan bool)
	account := &shared{mutex: &sync.Mutex{}}

	for i := 0; i < Conc; i++ {
		go adder(account, done)
	}

	// Wait for all goroutines to complete
	for i := 0; i < Conc; i++ {
		<-done
	}

	println(account.getCounter()) // Should == Conc x N
}
