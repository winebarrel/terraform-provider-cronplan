# terraform-provider-cronplan

[![test](https://github.com/winebarrel/terraform-provider-cronplan/actions/workflows/test.yml/badge.svg)](https://github.com/winebarrel/terraform-provider-cronplan/actions/workflows/test.yml)
[![terraform docs](https://img.shields.io/badge/terraform-docs-%35835CC?logo=terraform)](https://registry.terraform.io/providers/winebarrel/cronplan/latest/docs)

Terraform provider to validate Amazon EventBridge Cron expressions.

If the cron expression is incorrect, an error will occur in terraform plan.

![](https://github.com/winebarrel/terraform-provider-cronplan/assets/117768/36f3a5db-32b9-4ebb-a554-2b67ebc85174)

## Usage

```tf
provider "cronplan" {
}

data "cronplan_expr" "every_weekday" {
  # NOTE: Expression is validated with terraform plan.
  expr = "cron(5 0 ? * MON-FRI *)"
  from = "2023-09-30 10:00:00 UTC" # Optional
}

output "every_weekday" {
  value = data.cronplan_expr.every_weekday.schedules
}

check "every_weekday_schedules" {
  assert {
    condition = data.cronplan_expr.every_weekday.schedules == tolist([
      "Mon, 02 Oct 2023 00:05:00",
      "Tue, 03 Oct 2023 00:05:00",
      "Wed, 04 Oct 2023 00:05:00",
      "Thu, 05 Oct 2023 00:05:00",
      "Fri, 06 Oct 2023 00:05:00",
      "Mon, 09 Oct 2023 00:05:00",
      "Tue, 10 Oct 2023 00:05:00",
      "Wed, 11 Oct 2023 00:05:00",
      "Thu, 12 Oct 2023 00:05:00",
      "Fri, 13 Oct 2023 00:05:00",
    ])

    error_message = "Unexpected schedule: \n${join("\n", data.cronplan_expr.every_weekday.schedules)}"
  }
}
```

## Run locally for development

```sh
cp cronplan.tf.sample cronplan.tf
make tf-plan
```

## Related Links

* https://github.com/winebarrel/cronplan
* [Cron expressions reference - Amazon EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-cron-expressions.html)
