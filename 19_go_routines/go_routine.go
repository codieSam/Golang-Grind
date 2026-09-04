package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done() // this will tell the waitgroup that this go routine is done and it can decrement the counter by 1
	fmt.Println("Doing task", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i, &wg) //it will run all 11 functions paralelly and it will not wait for the function to complete

	}

	wg.Wait() // this will wait for all the go routines to complete before exiting the main function

	// time.Sleep(time.Second * 2) // instead of this we use waitgroup to wait for all the go routines to complete before exiting the main function

}
