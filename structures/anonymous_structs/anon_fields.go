package main

import "fmt"

type student struct {
	string
	int
}

func main() {
	st1 := student{
		"fawas",
		20,
	}

	fmt.Println(st1)
}
