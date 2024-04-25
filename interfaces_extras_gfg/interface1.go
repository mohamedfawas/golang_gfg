package main

import "fmt"

type tank interface {
	Area() float64
	Volume() float64
}

type myvalue struct {
	radius float64
	height float64
}

func (m myvalue) Area() float64 {
	return 2*m.radius*m.height + 2*3.14*m.radius*m.radius
}

func (m myvalue) Volume() float64 {
	return 3.14 * m.radius * m.radius * m.height
}

func main() {
	var t tank
	t = myvalue{10, 14}
	fmt.Println("area is : ", t.Area())
	fmt.Println("volume is : ", t.Volume())
}
