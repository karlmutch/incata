package incata

import (
	"github.com/satori/go.uuid"
)

// Reader interface for getting events
type Reader interface {
	Read(uuid.UUID) ([]Event, error)
}
