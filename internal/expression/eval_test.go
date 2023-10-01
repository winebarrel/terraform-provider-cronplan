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
	assert.ErrorContains(err, "Parse 'expr' failed: 'cron(0 0 ? * ? *)': '?' cannot be set to both day-of-month and day-of-week")
}
