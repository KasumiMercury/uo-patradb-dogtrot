package pulid

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

type ID string

var defaultEntropySource *ulid.MonotonicEntropy

func init() {
	defaultEntropySource = ulid.Monotonic(rand.Reader, 0)
}

func newULID() ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(time.Now()), defaultEntropySource)
}

func MustNew(prefix string) ID { return ID(prefix + fmt.Sprint(newULID())) }
