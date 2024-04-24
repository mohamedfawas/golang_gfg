// non struct types in receiver arguments

package main

import "fmt"

type data int // non struct type is defined

func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

func main() {
	val1 := data(10)
	val2 := data(2)
	result := val1.multiply(val2)
	fmt.Println("result of the given operation is : ", result)
}
