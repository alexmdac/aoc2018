package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	soln, err := run()
	require.NoError(t, err)
	require.Equal(t, &solution{}, soln)
}
