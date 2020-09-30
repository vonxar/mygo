package main

import "fmt"

func main() {
	i := 0
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	for index, val := range l {
		if index == 0 {
			i = val
			continue
		}
		if i >= val {
			i = val
		}
	}
	fmt.Println(i)
	i = 0
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	for _, val := range m {
		i += val
	}
	fmt.Println(i)
}
