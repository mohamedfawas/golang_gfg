package main

import "fmt"

func main() {
	myArray := []int{1, 2, 3}
	copyArray := make([]int, 0, len(myArray)) // length, capacity

	fmt.Println(myArray)
	fmt.Println(copyArray)

	copyArray = append(copyArray, myArray...)

	fmt.Println("==========after=============")
	fmt.Println(myArray)
	fmt.Println(copyArray)
}
