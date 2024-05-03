package record

import (
	"time"

	"github.com/carbans/simpledns/app/domain"
)

// Record represents a DNS record
type Record struct {
	Id        int
	Domain    *domain.Domain
	Type      string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewRecord creates a new record instance
func NewRecord(id int, domain *domain.Domain, recordType, value string) *Record {
	t := time.Now()
	createdAt := t
	updatedAt := t

	return &Record{
		Id:        id,
		Domain:    domain,
		Type:      recordType,
		Value:     value,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
