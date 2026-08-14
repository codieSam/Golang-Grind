package main

import "fmt"

func main() {
	// var name string = "Golang"
	var name = "Golang" //if you are assigning vakue at the time of declare then it will infer the type of the variable
	var isAudult = true
	var age int = 12
	//short_hand_syntax
	myName := "Short Hand Syntax" //useful when we have to assign and declare the varaible at a time !
	var price float32 = 50.5

	fmt.Println(name)
	fmt.Println(isAudult)
	fmt.Println(age)
	fmt.Println(myName)
	fmt.Println(price)
}
