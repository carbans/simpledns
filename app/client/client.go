package client

import (
	"fmt"

	"github.com/miekg/dns"
)

// Client is a DNS client
type DNSClient struct {
	upstream []string
}

func NewDNSClient(upstream []string) *DNSClient {
	return &DNSClient{
		upstream: upstream,
	}
}

func (c *DNSClient) Resolve(q dns.Question) (*dns.Msg, error) {
	m := new(dns.Msg)
	m.SetQuestion(q.Name, q.Qtype)
	m.RecursionDesired = true
	fmt.Println("Go resolve ", q.Name)
	var lastError error
	for _, upstream := range c.upstream {
		fmt.Println("Try upstream ", upstream)
		resp, err := dns.Exchange(m, upstream)
		if err != nil {
			lastError = err // Guarda el último error
			continue        // Intenta con el próximo upstream
		}
		return resp, nil // Si no hay error, devuelve la respuesta
	}

	// Si llegamos aquí, no se pudo resolver en ningún upstream, devuelve el último error
	return nil, lastError
}
