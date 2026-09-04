package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing task", id)
}

func main() {
	for i := 0; i < 10; i++ {
		go task(i) //it will run all 11 functions paralelly and it will not wait for the function to complete
	}

	time.Sleep(time.Second * 2)

}
