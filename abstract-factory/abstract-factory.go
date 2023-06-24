package abstract

import (
	"fmt"
)

type AuthFactory interface {
	CreateLogin() Auth
}

type Auth interface {
	Login()
}

type GoogleAuth struct {
	name string
}
type GoogleAuthFactory struct{}

type AppleAuth struct {
	name string
}
type AppleAuthFactory struct{}

// Implementation of interfaces
func (g *GoogleAuthFactory) CreateLogin() Auth {
	return &GoogleAuth{name: "Google"}
}

func (g *GoogleAuth) Login() {
	fmt.Printf("Logging in through: %s", g.name)
}

func (a *AppleAuthFactory) CreateLogin() Auth {
	return &AppleAuth{name: "Apple"}
}

func (a *AppleAuth) Login() {
	fmt.Printf("Logging in through: %s", a.name)
}

// returns Auth which has different implementations
// based on different providers
func GetProvider(factory AuthFactory) Auth {
	return factory.CreateLogin()
}
