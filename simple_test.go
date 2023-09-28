package practicum_module

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd_Zero(t *testing.T) {
	expected := 0

	actual := Add(0, 0)
	assert.Equal(t, expected, actual)
}
