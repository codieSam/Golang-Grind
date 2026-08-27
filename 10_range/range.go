package main

import "fmt"

//range is used for iterating over data structires

func main() {
	// nums := []int{6, 7, 8}
	// // for i := 0; i < len(nums); i++ {
	// // 	fmt.Println(nums[i])
	// // }

	// for i, num := range nums {

	// 	fmt.Println(num, i)
	// }

	// m := map[string]string{"name": "Golang", "area": "Backend"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	//unicode code points rune(rune = data structure for unicode code points)

	for k, v := range "Golang" {
		fmt.Println(k, string(v))
	}

}
