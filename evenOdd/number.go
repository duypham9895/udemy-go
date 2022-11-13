package main

import "fmt"

type number []int

func createNumbers(limit int) number {
	slices := number(make([]int, limit+1))

	for i := range slices {
		slices[i] = i
	}

	return slices
}

func (n number) validateEvenOdd() {
	for i := range n {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}
