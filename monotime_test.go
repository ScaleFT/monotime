package monotime

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNow(t *testing.T) {
	start := Now()
	require.NotEmpty(t, start)

	end := Now()
	require.NotEmpty(t, end)
	require.NotEmpty(t, start, end)

	d := Duration(start, end)
	require.NotEmpty(t, d)
}

func TestTimer(t *testing.T) {
	d := New()
	require.NotEmpty(t, d)

	x := d.Elapsed()
	require.NotEmpty(t, x)
	require.True(t, x > 0)
}
