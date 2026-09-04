package main

import "fmt"

// func printSlice[T int | string | bool](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func stringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// we even can use generics{}, comparable & like string | int | bool insteaad of any keywork in generics
type stack[T any] struct {
	elements []T
}

func main() {

	myStack := stack[int]{
		elements: []int{1, 2, 3},
	}

	fmt.Println(myStack)

	// nums := []int{1, 2, 3, 4, 5}
	// // names := []string{"Alice", "Bob", "Charlie"}
	// isTrue := []bool{true, false, true}
	// printSlice(isTrue)
	// // stringSlice(names)
}
