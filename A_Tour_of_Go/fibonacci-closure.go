package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var i int
	var n [3]int
	return func() int {
		defer func() { i++ }()
		if i < 2 {
			n[i] = i
			return i
		} else {
			n[2] = n[0] + n[1]
			n[0], n[1] = n[1], n[2]
			return n[2]
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
