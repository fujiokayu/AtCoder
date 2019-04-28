package main

import (
	"fmt"
)

func calkByShift() int {
	var N, arg int
	fmt.Scan(&N)

	var bitsum int
	for i := 0; i < N; i++ {
		fmt.Scan(&arg)
		bitsum = (arg | bitsum)
	}

	var count int
	for (bitsum & 1) == 0 {
		bitsum = bitsum >> 1
		count++
	}

	return count
}

func main() {
	fmt.Println(calkByShift())
}
