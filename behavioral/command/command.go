/*
	The Command Pattern encapsulates a request as an object, thereby
	letting you parameterize other objects with different requests,
	queue or log requests, and support undoable operations.
*/

package command

import (
	"fmt"
)

type Order interface {
	execute()
}

type Stock struct {
	name     string
	quantity int
	count    int
}

func (s *Stock) Buy() {
	fmt.Printf("Bought %d stocks of %s", s.quantity, s.name)
	s.count += 1
}

func (s *Stock) Sell() {
	fmt.Printf("Sold %d stocks of %s", s.quantity, s.name)
	s.count -= 1
}

// Implementation of Order interface
type BuyOrder struct {
	stock Stock
}

func (b *BuyOrder) execute() {
	b.stock.Buy()
}

type SellOrder struct {
	stock Stock
}

func (s *SellOrder) execute() {
	s.stock.Sell()
}

// this is a queue of orders
type Broker struct {
	order []Order
}

func (b *Broker) TakeOrder(o Order) {
	b.order = append(b.order, o)
}

// Invokes execute command
func (b *Broker) PlaceOrder() {
	for _, o := range b.order {
		o.execute()
	}

	b.order = make([]Order, 0)
}

func main() {
	// client creating command objects
	stock := Stock{name: "AMZN", quantity: 100, count: 0}
	buyO := &BuyOrder{stock}
	sellO := &SellOrder{stock}

	// set the commmands to the invoker
	broker := &Broker{}
	broker.TakeOrder(buyO)
	broker.TakeOrder(sellO)

	// invokes the command
	broker.PlaceOrder()
}
