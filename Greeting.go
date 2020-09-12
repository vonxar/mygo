package main

import (
	"fmt"
)

type Student struct {
	name string
}

func (s Student) CalAvg(data []float64) (avgResult float64) {
	sum := 0.0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	avgResult = sum / float64(len(data))
	return
}

func (s Student) Judge(avg float64) (judgeResult string) {
	if avg >= 60 {
		judgeResult = "passed"
	} else {
		judgeResult = "failed"
	}
	return
}

func main() {
	a001 := Student{name: "sato"}
	data := []float64{70, 65, 50, 90, 30}
	var avg float64 = a001.CalAvg(data)
	result := a001.Judge(avg)
	fmt.Println(avg)
	fmt.Println(a001.name + " " + result)
	fmt.Println("----------------------")
	a002 := Student{name: "dai"}
	data = []float64{55, 95, 90, 80, 20}
	avg = a002.CalAvg(data)
	result = a002.Judge(avg)
	fmt.Println(avg)
	fmt.Println(a002.name + " " + result)

}
