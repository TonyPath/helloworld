package helloworld

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_add(t *testing.T) {
	r := add(1, 1)
	require.Equal(t, 2, r)
}

func Test_mul(t *testing.T) {
	r := mul(1, 1)
	require.Equal(t, 1, r)
}
