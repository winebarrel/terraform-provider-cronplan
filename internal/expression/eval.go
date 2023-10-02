package expression

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/araddon/dateparse"
	cp "github.com/winebarrel/cronplan"
)

const (
	atExprFormat = "2006-01-02T15:04:05"
)

var (
	scheduleExprCron *regexp.Regexp
	scheduleExprRate *regexp.Regexp
	scheduleExprAt   *regexp.Regexp
	rateExprRegexp   *regexp.Regexp
)

func init() {
	scheduleExprCron = regexp.MustCompile(`^cron\(([^)]+)\)$`)
	scheduleExprRate = regexp.MustCompile(`^rate\(([^)]+)\)$`)
	scheduleExprAt = regexp.MustCompile(`^at\(([^)]+)\)$`)
	rateExprRegexp = regexp.MustCompile(`^(\d+) (?:minute|minutes|hour|hours|day|days)$`)
}

func Validate(expr string) error {
	_, err := Eval(expr, "", 0)
	return err
}

func Eval(expr string, fromStr string, n int) ([]string, error) {
	if m := scheduleExprCron.FindStringSubmatch(expr); len(m) == 2 {
		return evalCron(m[1], fromStr, n)
	} else if m := scheduleExprRate.FindStringSubmatch(expr); len(m) == 2 {
		return []string{}, evalRate(m[1])
	} else if m := scheduleExprAt.FindStringSubmatch(expr); len(m) == 2 {
		return []string{}, evalAt(m[1])
	} else {
		return nil, fmt.Errorf("Unsupported schedule expression: '%s'", expr)
	}
}

func evalCron(expr string, fromStr string, n int) ([]string, error) {
	cron, err := cp.Parse(expr)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse expr: 'cron(%s)': %w", expr, err)
	}

	from := time.Now()

	if fromStr != "" {
		from, err = dateparse.ParseAny(fromStr)

		if err != nil {
			return nil, fmt.Errorf("Failed to parse 'from': %w", err)
		}
	}

	schedule := []string{}
	next := cron.NextN(from, n)

	for _, v := range next {
		schedule = append(schedule, v.Format("Mon, 02 Jan 2006 15:04:05"))
	}

	return schedule, nil
}

func evalRate(expr string) error {
	m := rateExprRegexp.FindStringSubmatch(expr)

	if len(m) != 2 {
		return fmt.Errorf("Failed to parse expr: 'rate(%s)': does not match '%s'", expr, rateExprRegexp)
	}

	v, err := strconv.Atoi(m[1])

	if err != nil {
		return fmt.Errorf("Failed to parse rate value: 'rate(%s)': %w", expr, err)
	}

	if v < 1 {
		return fmt.Errorf("Rate expr value must be '>= 1': 'rate(%s)'", expr)
	}

	return nil
}

func evalAt(expr string) error {
	_, err := time.Parse(atExprFormat, expr)

	if err != nil {
		return fmt.Errorf("Failed to parse expr: 'at(%s)': does not match '%s'", expr, atExprFormat)
	}

	return nil
}
