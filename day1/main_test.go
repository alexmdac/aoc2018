package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	soln, err := run()
	require.NoError(t, err)
	require.Equal(t, solution{sum: 439, firstRepeatedSum: 124645}, soln)
}
