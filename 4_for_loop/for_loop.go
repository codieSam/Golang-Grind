package main

import (
	"fmt"
)

//for is only construct in go for looping

func main() {
	// while loop is also done by using for
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//infinite loop

	// for {
	// 	println("1")
	// }

	//classic for loop

	for j := 0; j <= 3; j++ {
		// break
		if j == 2 {
			continue
		}
		fmt.Println(j)

	}

	// range in new go 1.22

	for k := range 3 {
		fmt.Println(k)
	}

}
