package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch statement
	i := 3

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	//multiple conditions in a switch statement

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//Type switch statement

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("Its an boolean")
		default:
			fmt.Printf("I don't know what type is this %T\n", t)
		}
	}
	whoAmI(52.56)

}
