package helloworld

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_add(t *testing.T) {
	r := add(1, 1)
	require.Equal(t, 2, r)
}
