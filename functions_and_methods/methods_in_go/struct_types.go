package main

import "fmt"

//define the struct
type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

func (a author) show() {
	fmt.Println("author is : ", a.name)
	fmt.Println("branch is : ", a.branch)
	fmt.Println("particles : ", a.particles)
	fmt.Println("salary is : ", a.salary)
}

func main() {
	res := author{
		name:      "fawas",
		branch:    "mech",
		particles: 10,
		salary:    20000,
	}

	res.show()
}
