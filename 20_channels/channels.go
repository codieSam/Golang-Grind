package main

import "fmt"

//sending

// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second)
// 	}

// fmt.Println("Processing number", <-numChan)
// }

// receiving
func sum(result chan int, num1 int, num2 int) {
	sumResult := num1 + num2
	result <- sumResult
}

func main() {

	result := make(chan int)

	go sum(result, 4, 5)

	res := <-result

	fmt.Println("Sum is", res)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.IntN(100)
	// }

	// messageChan := make(chan string)

	// messageChan <- "Ping"

	// msg:= <- messageChan

	// fmt.Println(msg)

}
