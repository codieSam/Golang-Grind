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

//constructer function to create a new order

func newOrder(id string, amount float32, status string) *order {
	//initial steup goes here ...
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder

}

// receiver function to change the status of the order
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	// we can make a striuct inside main function if we want to use it only in main function

	language := struct {
		name   string
		isGood bool
	}{"Golang", true}

	fmt.Println(language)

	myOrder := newOrder("1", 30.50, "received")
	fmt.Println(myOrder)

	// myOrder1 := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "pending",
	// }
	// myOrder1.changeStatus("confirmed")
	// fmt.Println(myOrder1.status)
	// fmt.Println(myOrder1.getAmount())
	// myOrder2 := order{
	// 	id:     "2",
	// 	amount: 100.00,
	// 	status: "received",
	// }

	// myOrder1.createdAt = time.Now()
	// myOrder2.createdAt = time.Now()

	// fmt.Println("Order struct", myOrder1)
	// fmt.Println("Order struct", myOrder2)

}
