package adapter

import (
	"fmt"
)

type Sales interface {
	Buy()
	Sell()
}

type Leasing struct {
	name  string
	year  int
	price int
}

func (l *Leasing) Lease() {
	fmt.Printf("Leasing %d %s at  %d", l.year, l.name, l.price)
}

func (l *Leasing) Return() {
	fmt.Printf("Returning %d %s at %d", l.year, l.name, l.price)
}

// implements adapter for leased vehicle
type Adapter struct {
	lease Leasing
}

func (a *Adapter) Buy() {
	a.lease.Lease()
}

func (a *Adapter) Sell() {
	a.lease.Return()
}

// implements Sale for purchased vehicle
type Purchase struct {
	name  string
	year  int
	price int
}

func (p *Purchase) Buy() {
	fmt.Printf("Purchasing %d %s at  %d", p.year, p.name, p.price)
}

func (p *Purchase) Sell() {
	fmt.Printf("Selling %d %s at  %d", p.year, p.name, p.price)
}

// Demo
func main() {
	lease := Leasing{
		name:  "Porsche",
		year:  2022,
		price: 80000,
	}

	purchase := Purchase{
		name:  "BMW",
		year:  2022,
		price: 70000,
	}

	adapter := Adapter{lease}

	adapter.Buy()
	adapter.Sell()

	purchase.Buy()
	purchase.Sell()

}
