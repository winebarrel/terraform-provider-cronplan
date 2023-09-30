package cronplan_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestExpr_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers:  testProviders,
		Steps: []resource.TestStep{
			{
				Config: `
					data "cronplan_expr" "every_day" {
						cron = "0 0 * * ? *"
						from = "2022-03-14 11:12:30 UTC"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "cron", "0 0 * * ? *"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "from", "2022-03-14 11:12:30 UTC"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "num_schedules", "10"),
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
			{
				Config: `
					data "cronplan_expr" "every_day" {
						cron = "5 3 * * ? *"
						from = "2022-03-14 11:12:30 UTC"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "cron", "5 3 * * ? *"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "from", "2022-03-14 11:12:30 UTC"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "num_schedules", "10"),
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
			{
				Config: `
					data "cronplan_expr" "every_day" {
						cron = "5 3 * * ? *"
						from = "2022-04-15 12:22:30 UTC"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "cron", "5 3 * * ? *"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "from", "2022-04-15 12:22:30 UTC"),
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "num_schedules", "10"),
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
			{
				Config: `
					data "cronplan_expr" "every_day" {
						cron          = "5 3 * * ? *"
						from          = "2022-04-15 12:22:30 UTC"
						num_schedules = 3
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.cronplan_expr.every_day", "cron", "5 3 * * ? *"),
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

func TestExpr_validationError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		Providers:  testProviders,
		Steps: []resource.TestStep{
			{
				Config: `
					data "cronplan_expr" "every_day" {
						cron = "5 3 ? * ? *"
						from = "2022-04-15 12:22:30 UTC"
					}
				`,
				ExpectError: regexp.MustCompile(`'?' cannot be set to both day-of-month and day-of-week`),
				PlanOnly:    true,
			},
		},
	})
}
