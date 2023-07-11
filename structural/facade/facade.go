/*
	The Facade Pattern provides a unified interface to a set of interfaces
	in a subsystem. Facade defines a higher-level interface that makes
	the substystem easier to use
*/

package main

import (
	"fmt"
)

type Login struct{}

func (l *Login) sendRequest() {
	fmt.Println("sending authentication request...")
}

func (l *Login) queryPassword() {
	fmt.Println("querying password from DB...")
}

func (l *Login) validatePassword() {
	fmt.Println("validating password...")
}

func (l *Login) sendResponse() {
	fmt.Println("authetication successful!")
}

type Authenticate interface {
	authenticate()
}

// create facade to hide implementations of each step
func (l *Login) authenticate() {
	l.sendRequest()
	l.queryPassword()
	l.validatePassword()
	l.sendResponse()
}

func main() {
	l := Login{}
	l.authenticate()
}
