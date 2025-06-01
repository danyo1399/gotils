package gotils

import (
	"strings"
	"github.com/google/uuid"
)

// NewUUIDV7Id returns a new UUID string that is formatted as a UUID v7 string.
// UUID v7 is a time ordered UUID
func NewUUIDV7Id() string {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	return id.String()
}

func NewTimeOrderedId() string {
	return strings.ReplaceAll(NewUUIDV7Id(), "-", "")
}
