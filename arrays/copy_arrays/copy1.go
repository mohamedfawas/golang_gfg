package main

import "fmt"

func main() {
	arr1 := [5]string{"jude", "camavinga", "vini", "guler", "rudiger"}
	fmt.Println("the first array created is : ", arr1)
	arr2 := arr1 // pass by value
	fmt.Println("the second array created is : ", arr2)
	arr2[0] = "kepa"
	fmt.Println("=======after changing value at first index ======================")
	fmt.Println("the first array created is : ", arr1)
	fmt.Println("the second array created is : ", arr2)

	fmt.Println()

	arr3 := [5]string{"jude", "camavinga", "vini", "guler", "rudiger"}
	fmt.Println("the first array created is : ", arr3)
	arr4 := &arr3 // pass by reference
	fmt.Println("the second array created is : ", arr4)
	arr4[0] = "kepa"
	fmt.Println("=======after changing value at first index ======================")
	fmt.Println("the first array created is : ", arr3)
	fmt.Println("the second array created is : ", arr4)
}
