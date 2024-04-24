package main

import "fmt"

func main() {
	student := make(map[int]string)
	fmt.Println(student)
	student[1] = "Vinicius"
	student[2] = "Bellingham"
	fmt.Println(student)
	name, ok := student[1]
	fmt.Println(name, " : ", ok)

	// shorter way to create maps
	cricketers := map[int]string{
		1: "Virat",
		2: "Rohith",
		3: "Rahul",
	}

	fmt.Println(cricketers)
}
