package domain

import (
	"errors"
	"fmt"
	"time"
)

var domains = []*Domain{
	NewDomain(1, "example.com"),
	NewDomain(2, "example.net"),
	NewDomain(3, "example.org"),
}

// Domain represents a domain name
type Domain struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewDomain creates a new domain instance
func NewDomain(id int, name string) *Domain {

	t := time.Now()
	createdAt := t
	updatedAt := t

	return &Domain{
		Id:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

// GetDomains returns a list of domain names
func GetDomains() []*Domain {
	return domains
}

// GetDomainByName returns a domain by name
func GetDomainByName(name string) (*Domain, error) {
	for _, d := range domains {
		fmt.Println(d.Name)
		if d.Name == name {
			fmt.Println("Found domain")
			return d, nil
		}
	}

	return nil, errors.New("Domain not found")
}
