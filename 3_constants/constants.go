package main

import "fmt"

const PI = 13.77 //we can declare the var/constants outside of the function but we can't use short_hand_syntax

func main() {
	const name string = "Golang" //as this is the constant this cannot be reassigned !
	fmt.Println(name)
	fmt.Println(PI)

	const ( //constant can be declared in group also
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)
}
