/*
	The Factory Pattern defines an interface for creating an object,
	but lets subclasses decide which class to instantiate.
	it lets a class defer instantiation to subclasses.
*/

package factory

import (
	"fmt"
)

type Login interface {
	Login()
}

type GoogleAuth struct {
	name string
}

func NewGoogleAuth() GoogleAuth {
	return GoogleAuth{name: "Google"}
}

// Implementation of Login for Google
func (g GoogleAuth) Login() {
	fmt.Printf("Logging in with: %s", g.name)
}

type AppleAuth struct {
	name string
}

func NewAppleAuth() AppleAuth {
	return AppleAuth{name: "Apple"}
}

// Implementation of Login for Apple
func (a AppleAuth) Login() {
	fmt.Printf("Logging in with: %s", a.name)
}

// Factory
func GetProvider(provider string) (Login, error) {
	switch {
	case provider == "Google":
		return NewGoogleAuth(), nil
	case provider == "Apple":
		return NewAppleAuth(), nil
	default:
		return nil, fmt.Errorf("provider type %s not recognized", provider)
	}
}
