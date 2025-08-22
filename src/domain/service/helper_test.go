package service

import (
	"assessment/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildEvents(t *testing.T) {
	plannings := []entity.Planning{
		{
			StartAt:  mustParseTime(t, "2019-05-30 00:00:00"),
			EndAt:    mustParseTime(t, "2019-06-01 00:00:00"),
			Quantity: 2,
		},
		{
			StartAt:  mustParseTime(t, "2019-05-31 00:00:00"),
			EndAt:    mustParseTime(t, "2019-06-02 00:00:00"),
			Quantity: 3,
		},
	}

	events := buildEvents(plannings)

	expected := []event{
		{t: mustParseTime(t, "2019-05-30 00:00:00"), change: 2},
		{t: mustParseTime(t, "2019-05-31 00:00:00"), change: 3},
		{t: mustParseTime(t, "2019-06-01 00:00:00"), change: -2},
		{t: mustParseTime(t, "2019-06-02 00:00:00"), change: -3},
	}

	assert.Equal(t, expected, events, "events should be built and sorted correctly")
}

func TestMaxConcurrentUsage(t *testing.T) {
	events := []event{
		{t: mustParseTime(t, "2019-05-30 00:00:00"), change: 2},
		{t: mustParseTime(t, "2019-05-31 00:00:00"), change: 3},
		{t: mustParseTime(t, "2019-06-01 00:00:00"), change: -2},
		{t: mustParseTime(t, "2019-06-02 00:00:00"), change: -3},
	}

	got := maxConcurrentUsage(events)
	assert.Equal(t, int64(5), got, "maxConcurrentUsage should calculate max correctly")
}

func TestBuildEvents_Empty(t *testing.T) {
	events := buildEvents(nil)
	assert.Empty(t, events, "no plannings should produce no events")
}

func TestMaxConcurrentUsage_Empty(t *testing.T) {
	got := maxConcurrentUsage(nil)
	assert.Equal(t, int64(0), got, "empty events should yield max usage 0")
}
func mustParseTime(t *testing.T, s string) time.Time {
	t.Helper()
	parsed, err := time.Parse("2006-01-02 15:04:05", s)
	require.NoError(t, err, "failed to parse time %s", s)
	return parsed
}