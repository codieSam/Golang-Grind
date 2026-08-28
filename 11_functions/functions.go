package main

import "fmt"

func myMessage(fName string, Age int) {
	fmt.Println("Hello", fName, "you are", Age, "years old")
}

func main() {
	myMessage("John", 25)
}
