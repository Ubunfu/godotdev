package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

/*
Run this with -gcflags="-m"
I'm not surprised to see that the params passed to MakePersonPointer
escape the stack, because this fxn returns a pointer, which
disqualifies them from eligibility for storage on the stack.
*/
func main() {
	_ = MakePerson("PJ", "Allen", 3)
	_ = MakePersonPointer("Squid", "Allen", 2)
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}
