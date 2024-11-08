package provider_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestExpr_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1
			{
				Config: `
					data "cronplan_expr" "every_day" {
						expr = "cron(0 0 * * ? *)"
						from = "2022-03-14 11:12:30 UTC"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "expr", "cron(0 0 * * ? *)"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "from", "2022-03-14 11:12:30 UTC"),
					resource.TestCheckNoResourceAttr("data.cronplan_expr.every_day", "num_schedules"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.#", "10"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.0", "Tue, 15 Mar 2022 00:00:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.1", "Wed, 16 Mar 2022 00:00:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.2", "Thu, 17 Mar 2022 00:00:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.3", "Fri, 18 Mar 2022 00:00:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.4", "Sat, 19 Mar 2022 00:00:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.5", "Sun, 20 Mar 2022 00:00:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.6", "Mon, 21 Mar 2022 00:00:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.7", "Tue, 22 Mar 2022 00:00:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.8", "Wed, 23 Mar 2022 00:00:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.9", "Thu, 24 Mar 2022 00:00:00"),
				),
			},
			// Step 2
			{
				Config: `
					data "cronplan_expr" "every_day" {
						expr = "cron(5 3 * * ? *)"
						from = "2022-03-14 11:12:30 UTC"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "expr", "cron(5 3 * * ? *)"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "from", "2022-03-14 11:12:30 UTC"),
					resource.TestCheckNoResourceAttr("data.cronplan_expr.every_day", "num_schedules"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.#", "10"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.0", "Tue, 15 Mar 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.1", "Wed, 16 Mar 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.2", "Thu, 17 Mar 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.3", "Fri, 18 Mar 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.4", "Sat, 19 Mar 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.5", "Sun, 20 Mar 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.6", "Mon, 21 Mar 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.7", "Tue, 22 Mar 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.8", "Wed, 23 Mar 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.9", "Thu, 24 Mar 2022 03:05:00"),
				),
			},
			// Step 3
			{
				Config: `
					data "cronplan_expr" "every_day" {
						expr = "cron(5 3 * * ? *)"
						from = "2022-04-15 12:22:30 UTC"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "expr", "cron(5 3 * * ? *)"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "from", "2022-04-15 12:22:30 UTC"),
					resource.TestCheckNoResourceAttr("data.cronplan_expr.every_day", "num_schedules"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.#", "10"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.0", "Sat, 16 Apr 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.1", "Sun, 17 Apr 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.2", "Mon, 18 Apr 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.3", "Tue, 19 Apr 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.4", "Wed, 20 Apr 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.5", "Thu, 21 Apr 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.6", "Fri, 22 Apr 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.7", "Sat, 23 Apr 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.8", "Sun, 24 Apr 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.9", "Mon, 25 Apr 2022 03:05:00"),
				),
			},
			// Step 4
			{
				Config: `
					data "cronplan_expr" "every_day" {
						expr          = "cron(5 3 * * ? *)"
						from          = "2022-04-15 12:22:30 UTC"
						num_schedules = 3
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "expr", "cron(5 3 * * ? *)"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "from", "2022-04-15 12:22:30 UTC"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "num_schedules", "3"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.#", "3"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.0", "Sat, 16 Apr 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.1", "Sun, 17 Apr 2022 03:05:00"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "schedules.2", "Mon, 18 Apr 2022 03:05:00"),
				),
			},
		},
	})
}

func TestExpr_rate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "cronplan_expr" "every_minutes" {
						expr = "rate(1 minute)"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cronplan_expr.every_minutes", "expr", "rate(1 minute)"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_minutes", "schedules.#", "0"),
				),
			},
		},
	})
}

func TestExpr_rateErr(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "cronplan_expr" "every_minutes" {
						expr = "rate(minute)"
					}
				`,
				ExpectError: regexp.MustCompile(strings.ReplaceAll(` `, `\s+`, regexp.QuoteMeta(`Failed to parse expr: 'rate(minute)': does not match '^(\d+) (?:minute|minutes|hour|hours|day|days)$'`))),
				PlanOnly:    true,
			},
			{
				Config: `
					data "cronplan_expr" "every_minutes" {
						expr = "rate(0 minute)"
					}
				`,
				ExpectError: regexp.MustCompile(strings.ReplaceAll(` `, `\s+`, regexp.QuoteMeta("Rate expr value must be '>= 1': 'rate(0 minute)'"))),
				PlanOnly:    true,
			},
		},
	})
}

func TestExpr_at(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "cronplan_expr" "at_expr" {
						expr = "at(2016-03-04T17:27:00)"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cronplan_expr.at_expr", "expr", "at(2016-03-04T17:27:00)"),
					resource.TestCheckResourceAttr("data.cronplan_expr.at_expr", "schedules.#", "0"),
				),
			},
		},
	})
}

func TestExpr_atErr(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "cronplan_expr" "at_expr" {
						expr = "at(2016/03/04T17:27:00)"
					}
				`,
				ExpectError: regexp.MustCompile(strings.ReplaceAll(` `, `\s+`, regexp.QuoteMeta("Failed to parse expr: 'at(2016/03/04T17:27:00)': does not match '2006-01-02T15:04:05'"))),
				PlanOnly:    true,
			},
		},
	})
}

func TestExpr_validationError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "cronplan_expr" "every_day" {
						expr = "cron(5 3 ? * ? *)"
						from = "2022-04-15 12:22:30 UTC"
					}
				`,
				ExpectError: regexp.MustCompile(strings.ReplaceAll(` `, `\s+`, regexp.QuoteMeta(`Failed to parse expr: 'cron(5 3 ? * ? *)': '?' cannot be set to both day-of-month and day-of-week`))),
				PlanOnly:    true,
			},
			{
				Config: `
					data "cronplan_expr" "every_day" {
						expr = "cron(5 3 * * ? *)"
						from = "London Bridge is broken down"
					}
				`,
				ExpectError: regexp.MustCompile(strings.ReplaceAll(` `, `\s+`, regexp.QuoteMeta(`Failed to parse 'from': Could not find format for "London Bridge is broken down"`))),
				PlanOnly:    true,
			},
		},
	})
}

func TestExpr_unsupported(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "cronplan_expr" "every_day" {
						expr = "norc(5 3 ? * ? *)"
						from = "2022-04-15 12:22:30 UTC"
					}
				`,
				ExpectError: regexp.MustCompile(strings.ReplaceAll(` `, `\s+`, regexp.QuoteMeta(`Unsupported schedule expression: 'norc(5 3 ? * ? *)'`))),
				PlanOnly:    true,
			},
		},
	})
}
