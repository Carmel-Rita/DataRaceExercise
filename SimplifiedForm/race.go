package main

import (
	"fmt"
)

const (
	N    = 1000
	Conc = 1000
)

func add(sum int) int {

	for i := 0; i < N; i++ {
		sum += 1

	}
	return sum
}

func main() {

	sum := 0
	total := 0
	for i := 0; i < Conc; i++ {
		total += add(sum)

	}
	fmt.Println("total = ", total)

}
