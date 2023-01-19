package ptr_test

import (
	"testing"
	"time"

	"github.com/candiduslynx/ptr"
	"github.com/stretchr/testify/require"
)

func TestFrom(t *testing.T) {
	require.IsType(t, 0, ptr.From(new(int)))
	require.IsType(t, int32(0), ptr.From(new(int32)))
	require.IsType(t, int64(0), ptr.From(new(int64)))
	require.IsType(t, struct{}{}, ptr.From(new(struct{})))
	require.IsType(t, time.Time{}, ptr.From(new(time.Time)))
	require.IsType(t, new(time.Time), ptr.From(new(*time.Time)))
}

func TestFrom_Panics(t *testing.T) {
	require.Panics(t, func() {
		ptr.From((*float32)(nil))
	})
}

func TestFrom_NotPanics(t *testing.T) {
	require.NotPanics(t, func() {
		ptr.From((*float32)(nil), 0)
	})
}

func TestTo(t *testing.T) {
	require.IsType(t, new(int), ptr.To(0))
	require.IsType(t, new(int32), ptr.To(int32(0)))
	require.IsType(t, new(int64), ptr.To(int64(0)))
	require.IsType(t, new(struct{}), ptr.To(struct{}{}))
	require.IsType(t, new(time.Time), ptr.To(time.Time{}))
	require.IsType(t, new(*time.Time), ptr.To(new(time.Time)))
}
