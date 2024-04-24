package main

import "fmt"

func main() {
	person := struct {
		name   string
		branch string
	}{
		name:   "fawas",
		branch: "cse",
	}

	fmt.Println("the created anon struct values are : ", person)
}
