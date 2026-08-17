package main

import "fmt"

func main() {
	// age := 12

	// if age:=18; age >= 18 {  --->> We even can declare the variable in the if statement
	// 	fmt.Println("Person is an adult.")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	var role = "admin"
	var hasPermission = true

	if role == "admin" || hasPermission {
		fmt.Println("Access granted")
	}
}
