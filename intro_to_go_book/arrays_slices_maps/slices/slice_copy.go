package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	fmt.Println(slice1, slice2)
	copy(slice2, slice1) // copies first two values in the slice1 to slice2, since slice2 have only two
	fmt.Println(slice2)

}
