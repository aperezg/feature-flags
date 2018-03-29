package identity

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

// NewID Return an new ID regardless of the library used for that purpose
func NewID() string {
	t := time.Unix(1000000, 0)
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
