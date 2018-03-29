package identity

import "github.com/satori/go.uuid"

// NewID Return an new ID regardless of the library used for that purpose
func NewID() string {
	return string(uuid.NewV4().String())
}
