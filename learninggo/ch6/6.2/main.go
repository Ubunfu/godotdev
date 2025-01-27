package main

import "fmt"

func main() {
	cats := []string{"PJ", "Squid"}
	fmt.Printf("Cats before updating: %v\n", cats)
	UpdateSlice(cats, "Ryan")
	fmt.Printf("Cats before growing: %v\n", cats)
	GrowSlice(cats, "Allen")
	fmt.Printf("Cats after growing: %v\n", cats)
}

/*
Data within slices passed into functions can be modified in-place and persist outside of function block because slices are implemented with
pointers.  The data passed by value into the function is a pointer to the slice instead of the slice itself.  The slice within the function block
points the same memory.
*/
func UpdateSlice(strings []string, string string) {
	lastSlicePosition := len(strings) - 1
	strings[lastSlicePosition] = string
	fmt.Printf("Updated Slice: %v\n", strings)
}

/*
Slices grown inside a function block won't appear grown outside the function block because even though the slices inside and outside the function
block point to the same memory the slice structs within each block are independently tracking the length and capacity of the slice.  When the
slice grows within the function block, appended data is written to memory but the slice outside the function block isn't made aware of it.
*/
func GrowSlice(strings []string, string string) {
	strings = append(strings, string)
	fmt.Printf("Grew Slice: %v\n", strings)
}
