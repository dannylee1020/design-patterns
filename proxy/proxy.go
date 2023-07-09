package proxy

import (
	"fmt"
)

type API interface {
	SendResponse() error
}

type Client struct {
	ipAddress string
}

type Server struct {
	blockedIP []string
}

func (s *Server) SendResponse(c *Client) {
	fmt.Printf("sending response to client's ip address: %s", c.ipAddress)
}

type Proxy struct {
	Server
}

func (p *Proxy) SendResponse(c *Client) error {
	client := Client{"test-ip-address"}

	for _, ip := range p.Server.blockedIP {
		if ip == c.ipAddress {
			return fmt.Errorf("client IP is blocked")
		}
	}

	p.Server.SendResponse(&client)
	return nil
}
