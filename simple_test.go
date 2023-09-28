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

func TestAdd_One(t *testing.T) {
	expected := 1
	actual := Add(1, 0)

	assert.Equalf(t, expected, actual, "Not equal to 1: %d != %d", expected, actual)
}

func TestAdd_Svn(t *testing.T) {
	expected := 7
	actual := Add(3, 4)

	assert.Equal(t, expected, actual)
}

func TestAdd_MinusToZero(t *testing.T) {
	expected := 0
	actual := Add(-4, 4)

	assert.Equal(t, expected, actual)
}

func TestAdd_ToMinus(t *testing.T) {
	expected := -10
	actual := Add(-10, 0)

	assert.Equal(t, expected, actual)
}
