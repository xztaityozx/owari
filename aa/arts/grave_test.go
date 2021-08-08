package arts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGrave(t *testing.T) {
	g := NewGrave("text")
	assert.Equal(t, "text", g.text)
}
