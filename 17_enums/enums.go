package main

import "fmt"

//enumerated types

type OrderStatus string

const (
	Received   OrderStatus = "Received"
	Processing OrderStatus = "Processing"
	Shipped    OrderStatus = "Shipped"
	Delivered  OrderStatus = "Delivered"
	Cancelled  OrderStatus = "Cancelled"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	changeOrderStatus(Received)
	changeOrderStatus(Processing)
	changeOrderStatus(Shipped)
	changeOrderStatus(Delivered)
	changeOrderStatus(Cancelled)
}
