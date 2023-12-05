package pulid

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

var defaultEntropySource *ulid.MonotonicEntropy

func init() {
	defaultEntropySource = ulid.Monotonic(rand.Reader, 0)
}

func newULID() string {
	return ulid.MustNew(ulid.Timestamp(time.Now()), defaultEntropySource).String()
}

func MustNew(prefix string) string { return prefix + fmt.Sprint(newULID()) }
