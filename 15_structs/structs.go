package main

import (
	"fmt"
	"time"
)

// order structs

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

func main() {

	myOrder1 := order{
		id:     "1",
		amount: 50.00,
		status: "pending",
	}
	myOrder2 := order{
		id:     "2",
		amount: 100.00,
		status: "received",
	}

	myOrder1.createdAt = time.Now()
	myOrder2.createdAt = time.Now()

	fmt.Println("Order struct", myOrder1)
	fmt.Println("Order struct", myOrder2)

}
