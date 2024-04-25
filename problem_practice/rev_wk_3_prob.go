package main

import "fmt"

func main() {
	var nums []int
	for i := 1; i <= 100; i++ {
		if i%7 == 0 {
			nums = append(nums, i)
		}
	}
	//fmt.Println(nums)
	sq_map := make(map[int]int)
	for _, v := range nums {
		sq_map[v] = v * v
	}
	fmt.Println(sq_map)
}
