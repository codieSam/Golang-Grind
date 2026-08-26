package main

import "fmt"

//maps -> hash, object, dict

func main() {
	//creating map

	m := make(map[string]string)

	// setting an elements
	m["name"] = "Golang"
	m["area"] = "Backend"
	//get an element

	fmt.Println(m["name"], m["area"], m["version"])

}
