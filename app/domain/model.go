package domain

import (
	"time"
)

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
