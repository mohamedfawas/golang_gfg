package main

import "fmt"

func main() {
	myArray := []int{1, 2, 3}
	copyArray := make([]int, len(myArray))

	fmt.Println(myArray)
	fmt.Println(copyArray)

	copy(copyArray, myArray)

	fmt.Println("==========after=============")
	fmt.Println(myArray)
	fmt.Println(copyArray)
}
