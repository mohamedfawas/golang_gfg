package main

import "fmt"

type inter1 interface {
	Area() float64
	//Volume() float64
}

type data struct {
	length float64
	breadth float64
}

func (d data) Area() float64  {
	area := d.length*d.breadth
	return area
}

func main() {
	var t inter1
	t = data{5,4}
	fmt.Println("area is : ",t.Area()  )
}