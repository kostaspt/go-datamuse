package datamuse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatamuse_URL(t *testing.T) {
	assert.Equal(t, "https://api.datamuse.com", New().URL())
}
