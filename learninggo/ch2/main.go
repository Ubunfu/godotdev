package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
}

func ex1() {
	fmt.Println("=== EXERCISE 1 ===")

	var i int = 20
	var f float64 = float64(i)

	fmt.Println(i)
	fmt.Println(f)
}

func ex2() {
	fmt.Println("=== EXERCISE 2 ===")

	const value = 1

	var i int = value
	var f float64 = value

	fmt.Println(i)
	fmt.Println(f)
}

func ex3() {
	fmt.Println("=== EXERCISE 3 ===")

	var b byte = 255
	var smallI int32 = 2_147_483_647
	var bigI uint64 = 18_446_744_073_709_551_615

	b += 1
	smallI += 1
	bigI += 1

	println(b)
	println(smallI)
	println(bigI)
}
