package main

/*
TARGET_HEAP = CURR_HEAP + CURR_HEAP*GOGC/100
Ex. 8MB starting live heap

18MB = 8MB + (8MB+1MB+1MB)*100/100
*/

/*
Set GODEBUG=gctrace=1

GOGC: 50
~44 GCs
~3.2s

GOGC: 100
~20 GCs
~2.6s

GOGC: 200
~11 GCs
~2.1s

Pre-allocating the capacity makes it a lot faster.
*/
func main() {
	hugeSlice := []string{}
	// hugeSlice := make([]string, 0, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		hugeSlice = append(hugeSlice, "Ryan")
	}
}
