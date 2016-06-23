package main

import (
	"fmt"
)

const (
	N    = 1000
	Conc = 1000
)

var sum int = 0

func add(acc chan int) {

	for i := 0; i < N; i++ {
		sum += 1

	}
	acc <- sum

}

func main() {

	acc := make(chan int)
	total := 0
	for i := 0; i < Conc; i++ {
		go add(acc)
		total = <-acc
	}
	fmt.Println("total = ", total)

}
