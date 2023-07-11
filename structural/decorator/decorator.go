/*
	The Decorator Pattern attaches additional responsibilities to
	an object dynamically. Decorators provide a flexible alternative to
	subclassing for extending functionality.
*/

package decorator

import (
	"fmt"
)

type Component struct {
	msg string
}

func (c *Component) Print() string {
	return c.msg
}

func (c *Component) PrintDecorator(f func() string) {
	message := f()
	fmt.Println("Adding decoration to the original message - ", message)
}

func main() {
	c := Component{msg: "Hello World!"}
	c.PrintDecorator(c.Print)
}
