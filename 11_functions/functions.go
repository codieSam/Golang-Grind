package main

import "fmt"

// func myMessage(fName string, Age int) {
// 	fmt.Println("Hello", fName, "you are", Age, "years old")
// }

// func main() {
// 	myMessage("John", 25)
// }

// Return keyword is used to return a value from a function

func myMessage(fname string, Age int) string {
	return ("Hello " + fname + ", you are " + fmt.Sprintf("%d", Age) + " years old.")
}

func main() {
	fmt.Println(myMessage("John", 25))
}
