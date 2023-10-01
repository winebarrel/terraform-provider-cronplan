package expression

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/araddon/dateparse"
	cp "github.com/winebarrel/cronplan"
)

var (
	scheduleExprCron *regexp.Regexp
	scheduleExprRate *regexp.Regexp
	scheduleExprAt   *regexp.Regexp
	rateExpr         *regexp.Regexp
)

func init() {
	scheduleExprCron = regexp.MustCompile(`^cron\(([^)]+)\)$`)
	scheduleExprRate = regexp.MustCompile(`^rate\(([^)]+)\)$`)
	scheduleExprAt = regexp.MustCompile(`^at\([^)]+\)$`)
	rateExpr = regexp.MustCompile(`^(\d+) (?:minute|minutes|hour|hours|day|days)$`)
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
	} else if scheduleExprAt.MatchString(expr) {
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

func evalRate(expr string) error {
	m := rateExpr.FindStringSubmatch(expr)

	if len(m) != 2 {
		return fmt.Errorf("Unexpected rate expr: 'rate(%s)': does not match '%s'", expr, rateExpr)
	}

	v, err := strconv.Atoi(m[1])

	if err != nil {
		return fmt.Errorf("Parse expr value failed: 'rate(%s)': %w", expr, err)
	}

	if v < 1 {
		return fmt.Errorf("Expr value is less than or equal to 0: 'rate(%s)'", expr)
	}

	return nil
}
