package main

//maps -> hash, object, dict
import (
	"fmt"
)

func main() {
	// //creating map

	// m := make(map[string]string)

	// // setting an elements
	// m["name"] = "Golang"
	// m["area"] = "Backend"
	// //get an element

	// fmt.Println(m["name"], m["area"], m["version"])

	// m := make(map[string]int)

	// m["Volvo"] = 2010
	// m["BMW"] = 2015
	// m["Ford"] = 2005
	// delete(m, "Ford")
	// // clear(m)
	// fmt.Println(len(m))
	// fmt.Println(n)

	// fmt.Println(m["Volvo"], m["BMW"], m["Ford"], m["name"])

	n := map[string]int{"Volvo": 2010, "BMW": 2015, "Ford": 2005}

	v, ok := n["Ford"]
	fmt.Println(v)
	if ok {
		fmt.Println("It's okay")
	} else {
		fmt.Println("It's not okay")
	}

}
