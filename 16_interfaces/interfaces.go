package main

import "fmt"

//we give method to the interface and implement it in the struct. This is called polymorphism
type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter // interface implementation
}

// Open close principle: Open for extension, closed for modification

func (p payment) makePayment(amount float32) {
	// razorpayPayment := razorpay{} // interface implementation
	// stripePayment := stripe{}
	p.gateway.pay(amount)

	// razorpayPayment.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("Making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	// logic to make payment
	fmt.Println("Making payment using stripe", amount)
}

// if in future we want to add Paypal payment gateway then

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Making payment using Paypal", amount)
}

func (p paypal) refund(amount float32, account string) {

}

func main() {
	// stripePaymentGw := stripe{}
	paypalGw := paypal{}   //crreating an instance of paypal struct
	newPayment := payment{ //constructing a new payment struct and passing the paypal instance to the gateway field
		gateway: paypalGw,
	} //interface implementation
	newPayment.makePayment(100)
}
