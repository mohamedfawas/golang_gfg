package main

import "fmt"

type Author struct {
	name   string
	branch string
	year   int
}

type HR struct {
	details Author
}

func main() {
	result := HR{
		details: Author{"fawas", "mech", 2020},
	}
	fmt.Println("created struct is : ", result)
}
