package expression

import (
	"fmt"
	"regexp"
	"time"

	"github.com/araddon/dateparse"
	cp "github.com/winebarrel/cronplan"
)

var (
	ScheduleExprCron *regexp.Regexp
	ScheduleExprRate *regexp.Regexp
	ScheduleExprAt   *regexp.Regexp
)

func init() {
	ScheduleExprCron = regexp.MustCompile(`^cron\(([^)]+)\)$`)
	ScheduleExprRate = regexp.MustCompile(`^rate\([^)]+\)$`)
	ScheduleExprAt = regexp.MustCompile(`^at\([^)]+\)$`)
}

func Validate(expr string) error {
	_, err := Eval(expr, "", 0)
	return err
}

func Eval(expr string, fromStr string, n int) ([]string, error) {
	if m := ScheduleExprCron.FindStringSubmatch(expr); len(m) == 2 {
		return evalCron(m[1], fromStr, n)
	} else if ScheduleExprRate.MatchString(expr) {
		return []string{}, nil
	} else if ScheduleExprAt.MatchString(expr) {
		return []string{}, nil
	} else {
		return nil, fmt.Errorf("Unsupported schedule expression: '%s'", expr)
	}
}

func evalCron(expr string, fromStr string, n int) ([]string, error) {
	cron, err := cp.Parse(expr)

	if err != nil {
		return nil, fmt.Errorf("Parse 'expr' failed: 'cron(%s)': %w", expr, err)
	}

	from := time.Now()

	if fromStr != "" {
		from, err = dateparse.ParseAny(fromStr)

		if err != nil {
			return nil, fmt.Errorf("Parse 'from' failed: %w", err)
		}
	}

	schedule := []string{}
	next := cron.NextN(from, n)

	for _, v := range next {
		schedule = append(schedule, v.Format("Mon, 02 Jan 2006 15:04:05"))
	}

	return schedule, nil
}
