package main

import (
	"fmt"
	"sort"
)

func main() {
	mySlice := []int{11, 23, 53, 1, 3, 65, 23, 5, 74, 11, 33, 4, 3, 7}
	sort.Ints(mySlice)
	fmt.Println("given slice after sorting : ", mySlice)
}
