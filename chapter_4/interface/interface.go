package main

import "fmt"

type Card interface {
	Pay(int) bool
}

type ImposterCard interface {
	Pay(int) bool
}

type MasterCard struct {
	amount int
	userId int
}

type VisaCard struct {
	amount int
	userId int
}

func (MasterCard) Pay(value int) bool {
	fmt.Println("Paid Via Mastercard")
	return true
}

func (VisaCard) Pay(value int) bool {
	fmt.Println("Paid Via VisaCard")
	return true
}

func MakePayment(c Card, ic ImposterCard) {
	c.Pay(10)
	ic.Pay(10)
}

func main() {

	mc := MasterCard{amount: 100, userId: 1}
	vc := VisaCard{amount: 2000, userId: 2}

	MakePayment(mc, vc)
	MakePayment(vc, mc)
}
