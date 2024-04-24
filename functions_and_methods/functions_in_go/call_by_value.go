package main

import "fmt"

func swap(a, b int) (int, int) {
	var temp int
	temp = a
	a = b
	b = temp
	return a, b
}

func main() {
	var a, b int
	fmt.Println("give value for a : ")
	fmt.Scan(&a)
	fmt.Println("give value for b : ")
	fmt.Scan(&b)
	fmt.Println("before swap")
	fmt.Printf(" value of a is %d and b is %d \n", a, b)
	fmt.Println("after swap")
	swap(a, b) // call by value
	fmt.Printf("value of a is %d and b is %d ", a, b)
}
