package main

import (
	"fmt"
	"strings"
)

func joinStrings(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {
	fmt.Println(joinStrings()) // zero arguments
	fmt.Println(joinStrings("real madrid", "is", "a", "phenomenon"))
}
