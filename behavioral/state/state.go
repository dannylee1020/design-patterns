/*
	The State Pattern allows an object to alter its behavior when its
	internal state changes. The object will appear to change its class.
*/

package main

import (
	"fmt"
)

type State interface {
	Login()
	Logout()
}

type Context struct {
	authState bool
}

type Auth struct{}

func (a *Auth) Login(c *Context) {
	if c.authState {
		panic("User is already logged in")
	}
	fmt.Println("Logging in...")
	c.authState = true
}

func (a *Auth) Logout(c *Context) {
	if !c.authState {
		panic("User is not logged in")
	}

	fmt.Println("Logging out...")
	c.authState = false
}

func main() {
	ctx := Context{authState: false}
	auth := Auth{}

	auth.Login(&ctx)
	auth.Logout(&ctx)
	auth.Logout(&ctx)
}
