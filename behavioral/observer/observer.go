/*
	The Observer Pattern defines a one-to-many deopendency between objects
	so that when one object changes state, all of its dependents are notified
	and updated automatically.
*/

package observer

import (
	"fmt"
)

type Publisher interface {
	Register()
	Remove()
	Notify()
}

type Observer interface {
	Update()
}

// Implement publisher
type ConcretePublisher struct {
	observers []Observer
}

func NewPublisher() ConcretePublisher {
	return ConcretePublisher{}
}

func (c *ConcretePublisher) Register(obs Observer) {
	c.observers = append(c.observers, obs)
}

func (c *ConcretePublisher) Remove(obs Observer) {
	for i := 0; i < len(c.observers); i++ {
		if obs == c.observers[i] {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
		}
	}
}

func (c *ConcretePublisher) Notify() {
	fmt.Println("Got new message")

	for _, o := range c.observers {
		o.Update()
	}
}

// implement observer
type ConcreteObserver struct{}

func NewObserver() ConcreteObserver {
	return ConcreteObserver{}
}

func (o *ConcreteObserver) Update() {
	// update logic goes here
}
