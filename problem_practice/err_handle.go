package main

import (
	"errors"
	"fmt"
)

func show(i int) (int, error) {
	if i > 5 {
		return 10, nil
	} else {
		return 0, errors.New("error is happening")
	}
}

func main() {
	val, err := show(2)
	fmt.Println(val, err)
}
