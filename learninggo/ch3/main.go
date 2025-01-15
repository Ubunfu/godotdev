package main

import "fmt"

func main() {
	ex1()
	ex2()
    ex3()
}

func ex1() {
	fmt.Println("\n\nExample 1:\n")

	var greetings []string = []string{"Hello", "Hola", "नमस्ते", "こんにちは", "привет"}
	fmt.Printf("All greetings: %v\n", greetings)

	firstTwo := greetings[:2]
	secondToFourth := greetings[1:4]
	fourthAndFifth := greetings[3:]

	fmt.Printf("First two greetings: %v\n", firstTwo)
	fmt.Printf("Second to fourth greetings: %v\n", secondToFourth)
	fmt.Printf("Fourth and fifth greetings: %v\n", fourthAndFifth)

}

func ex2() {
	fmt.Println("\n\nExample 2:\n")

	var message string = "Hi 👧 and 👦"
	fmt.Println("Complete message: %v", message)

	fmt.Printf("Fourth rune: %c\n", []rune(message)[3])
}

func ex3() {
	fmt.Println("\n\nExample 3:\n")

	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	literalEmployee := Employee{"Ryan", "Allen", 1}
	fmt.Printf("Literal employee: %v\n", literalEmployee)

	structLiteralEmployee := Employee{
		firstName: "John",
		lastName:  "Wyatt",
		id:        2,
	}
	fmt.Printf("Struct Literal Employee: %v\n", structLiteralEmployee)

	var varDeclaredEmployee Employee = Employee{}
	varDeclaredEmployee.firstName = "Rakesh"
	varDeclaredEmployee.lastName = "Ray"
	varDeclaredEmployee.id = 3
	fmt.Printf("Var Declared Employee: %v\n", varDeclaredEmployee)
}
