package main

import "fmt"

func main() {
	{
		var i int = 10
		var p *int
		p = &i //ポイント
		var j int
		j = *p         // 中身
		fmt.Println(j) //10
	}

	var i int = 100
	var j int = 200
	var p1 *int
	var p2 *int
	p1 = &i //ポイント
	fmt.Println(p1)
	p2 = &j //ポイント
	fmt.Println(p2)
	i = *p1 + *p2 //中身足す 300
	fmt.Println(i)
	p2 = p1 //iポイント(100)
	fmt.Println(*p2)
	j = *p2 + i //300+300
	fmt.Println(j)
}
