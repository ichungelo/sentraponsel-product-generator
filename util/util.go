package util

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func GetUlid() (id string) {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	_id := ulid.MustNew(ulid.Timestamp(t), entropy)

	id = _id.String()

	return
}