package gotils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUUIDV7Id(t *testing.T) {
	id := NewUUIDV7Id()
	id2 := NewUUIDV7Id()
	assert.NotEmpty(t, id)
	assert.NotEqual(t, id, id2)
}

func TestNewTimeOrderedId(t *testing.T) {
	id := NewTimeOrderedId()
	id2 := NewTimeOrderedId()
	assert.NotEmpty(t, id)
	assert.NotEqual(t, id, id2)
}
