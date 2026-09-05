package main

import (
	"fmt"
	"time"
)

//sending

// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second)
// 	}

// fmt.Println("Processing number", <-numChan)
// }

// receiving
// func sum(result chan int, num1 int, num2 int) {
// 	sumResult := num1 + num2
// 	result <- sumResult
// }

// Go routine synchronizer

// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("processing task...")
// 	// done <- true
// }

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Sending email to ", email)
		time.Sleep(time.Second)
	}
}

func main() {

	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 0; i < 100; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("Done sending")
	<-done
	// emailChan <- "1@example.com"
	// emailChan <- "2@gmail.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// done := make(chan bool)

	// go task(done)

	// <-done //block the main function until the task is done and it will wait for the task to complete before exiting the main function

	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result

	// fmt.Println("Sum is", res)

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
