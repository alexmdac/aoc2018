package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	sum, rep, err := run()
	require.NoError(t, err)
	assert.Equal(t, 439, sum)
	assert.Equal(t, 124645, rep)
}
