package identity

import (
	"crypto/rand"

	"github.com/oklog/ulid"
)

// NewID Return an new ID regardless of the library used for that purpose
func NewID() string {
	return ulid.MustNew(ulid.Now(), rand.Reader).String()
}
