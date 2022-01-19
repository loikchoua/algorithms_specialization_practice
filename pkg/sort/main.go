package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/loikchoua/algorithms_practice_specialization/pkg/sort/merge"
)

func main() {
	// Generate random length for input
	len := rand.Intn(100)

	input := make([]int, len)

	// Insert random numbers into input array
	for i := 0; i < len; i++ {
		input[i] = rand.Int()
	}

	start := time.Now()
	fmt.Printf("initial array: %+v\n", input)

	// Sort the array
	input = merge.Sort(input)

	fmt.Printf("sorted array: %+v.\n elapsed time: %v", input, time.Since(start))
}
