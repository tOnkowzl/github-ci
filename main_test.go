package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum_ShouldReturnResultAsExpected(t *testing.T) {
	actual := sum(1, 1)

	assert.Equal(t, 2, actual)
}
