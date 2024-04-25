// How to get the current date and time with timestamp in local and other timezones ?
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Time at the location ", t.Location(), " is ", t)
	fmt.Println("weekday is : ", time.Now().Weekday())
	fmt.Println("year day is : ", time.Now().YearDay())
	fmt.Println("changing time location to EST ")
	loc, _ := time.LoadLocation("EST")
	fmt.Println("Current time in EST is : ", time.Now().In(loc))
}
