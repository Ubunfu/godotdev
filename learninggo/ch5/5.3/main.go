package main

import "fmt"

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // Should print "Hello Bob"
	fmt.Println(helloPrefix("Maria")) // Should print "Hello Maria"
}

func prefixer(prefix string) func(string) string {
	return func(s string) string {
		return prefix + " " + s
	}
}
