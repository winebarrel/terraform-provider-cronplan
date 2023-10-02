package expression_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/terraform-provider-cronplan/internal/expression"
)

func TestEvalCron_OK(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	schedules, err := expression.EvalCron("0 0 * * ? *", "2023/10/01 00:00", 3)
	require.NoError(err)
	assert.Equal([]string{
		"Sun, 01 Oct 2023 00:00:00",
		"Mon, 02 Oct 2023 00:00:00",
		"Tue, 03 Oct 2023 00:00:00",
	}, schedules)
}

func TestEvalCron_Err(t *testing.T) {
	assert := assert.New(t)

	_, err := expression.EvalCron("0 0 ? * ? *", "2023/10/01 00:00", 3)
	assert.ErrorContains(err, "Failed to parse expr: 'cron(0 0 ? * ? *)': '?' cannot be set to both day-of-month and day-of-week")
}

func TestEvalRate_OK(t *testing.T) {
	assert := assert.New(t)

	tt := []string{
		"1 minute",
		"1 minutes",
		"3 minute",
		"3 minutes",
		"1 hour",
		"1 hours",
		"3 hour",
		"3 hours",
		"1 day",
		"1 days",
		"3 day",
		"3 days",
		"01 minute",
		"01 minutes",
		"03 minute",
		"03 minutes",
		"01 hour",
		"01 hours",
		"03 hour",
		"03 hours",
		"01 day",
		"01 days",
		"03 day",
		"03 days",
	}

	for _, t := range tt {
		err := expression.EvalRate(t)
		assert.NoError(err)
	}
}

func TestEvalRate_Unexpected(t *testing.T) {
	assert := assert.New(t)

	tt := []string{
		"1  minute",
		"minutes",
		"3",
		"3 second",
	}

	for _, t := range tt {
		err := expression.EvalRate(t)
		assert.ErrorContains(err, "Failed to parse expr: 'rate("+t+")': does not match '^(\\d+) (?:minute|minutes|hour|hours|day|days)$'")
	}
}

func TestEvalRate_ZeroValue(t *testing.T) {
	assert := assert.New(t)
	err := expression.EvalRate("0 minute")
	assert.ErrorContains(err, "Rate expr value must be '>= 1': 'rate(0 minute)'")
}

func TestEvalAt_OK(t *testing.T) {
	assert := assert.New(t)
	err := expression.EvalAt("2023-10-01T01:02:03")
	assert.NoError(err)
}

func TestEvalAt_Err(t *testing.T) {
	assert := assert.New(t)
	err := expression.EvalAt("2023/10/01T01:02:03")
	assert.ErrorContains(err, "Failed to parse expr: 'at(2023/10/01T01:02:03)': does not match '2006-01-02T15:04:05'")
}
