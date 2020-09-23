package main

import(
	"fmt"
	"strconv"
)

func NumCheck(str string) bool {
	for _, r := range str {
		if '3' == r  {
			return true
		}
	}
	return false
}

func main() {
	var to int
for i := 0; i <= 30000; i++ {
	 if i%3 == 0 || NumCheck(strconv.Itoa(i)){
	 to += i%3
	 }
	 fmt.Println(to)

	}
	}