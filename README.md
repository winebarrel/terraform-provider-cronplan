# terraform-provider-cronplan

Terraform provider to validate Amazon EventBridge Cron expressions.

## Usage

```tf
provider "cronplan" {
}

data "cronplan_expr" "every_weekday" {
  cron = "5 0 ? * MON-FRI *" # NOTE: Cron expressions are validated with terraform plan
  from = "2023-09-30 10:00:00 UTC"
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
