package main

import (
	"fmt"
)

type Iterator interface {
	HasNext() bool
	CurrentItem() interface{}
}

type Printer struct {
	data  []string
	index int
}

func (p *Printer) HasNext() bool {
	p.index++
	return p.index < len(p.data)
}

func (p *Printer) CurrentItem() interface{} {
	return p.data[p.index]
}

// simple list that holds data

type List struct {
	data []string
}

func (l *List) Add(el string) {
	l.data = append(l.data, el)
}

func (l *List) CreateIterator() Iterator {
	newPrinter := &Printer{data: l.data, index: -1}
	return newPrinter
}

func main() {
	lst := List{data: make([]string, 0)}
	lst.Add("L")
	lst.Add("O")
	lst.Add("R")
	lst.Add("D")
	lst.Add("O")
	lst.Add("F")
	lst.Add("T")
	lst.Add("H")
	lst.Add("E")

	lstIterator := lst.CreateIterator()

	for lstIterator.HasNext() {
		fmt.Println("-", lstIterator.CurrentItem())
	}

}
