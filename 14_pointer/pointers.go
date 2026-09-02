package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // point to i

	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(*p) // read the updated value of i
	// fmt.Println(j)  // read j directly

	p = &j // point to j
	*p = *p / 37
	fmt.Println(j)

}
