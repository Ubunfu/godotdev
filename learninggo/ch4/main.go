package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomInts := ex1()
	ex2(randomInts)
	ex3()
}

func ex1() []int {
	fmt.Printf("\n\nExample 1: \n\n")

	// A slice of ints with initial length 0, and capacity 100
	// Prevents Go from having to re-size it
	randomInts := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		// Since length is 0, we can use append to populate it
		randomInts = append(randomInts, rand.Intn(100))
	}
	fmt.Printf("100 random integers: %v\n", randomInts)

	return randomInts
}

func ex2(randomInts []int) {
	fmt.Printf("\n\nExample 2: \n\n")

	for _, v := range randomInts {
		switch {
		case (v%2 == 0) && (v%3 == 0):
			fmt.Printf("%v is divisible by Six!\n", v)
		case v%2 == 0:
			fmt.Printf("%v is divisible by two!\n", v)
		case v%3 == 0:
			fmt.Printf("%v is divisible by three!\n", v)
		default:
			fmt.Println("Never mind.")
		}
	}
}

func ex3() {
	fmt.Printf("\n\nExample 3: \n\n")

	var total int
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
	fmt.Printf("Total is: %v\n", total)
    
    // There is likely a bug in this code
    // what it actually printed out is the nil-value of an 
    //      int (zero) plus the loop index, which is always 
    //      the same as the loop index.
    // What was likely intended is to add the loop index to 
    //      the current value, keeping state between iterations.
    // This bug is likely because `total` is being shadowed by 
    //      the for-loop.
}
