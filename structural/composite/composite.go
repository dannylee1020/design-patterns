package composite

import (
	"fmt"
)

type Component interface {
	Traverse()
}

type Leaf struct {
	value int
}

func (l *Leaf) Traverse() {
	fmt.Printf("%d", l.value)
}

type Composite struct {
	components []Component
}

func (c *Composite) Add(component Component) {
	c.components = append(c.components, component)
}

func (c *Composite) Remove(component Component) {
	for i, comp := range c.components {
		if comp == component {
			c.components = append(c.components[:i], c.components[i+1:]...)
			break
		}
	}
}

func (c *Composite) Traverse() {
	for i := 0; i < len(c.components); i++ {
		c.components[i].Traverse()
	}
}
