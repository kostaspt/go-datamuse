package datamuse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatamuse_Hyperlink(t *testing.T) {
	assert.Equal(t, "https://api.datamuse.com", New().Hyperlink())
}
