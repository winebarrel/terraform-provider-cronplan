output "every_weekday" {
  # If the expression is correct, the expression is returned.
  value = provider::cronplan::expr("cron(5 0 ? * MON-FRI *)")
  #=> value = "cron(5 0 ? * MON-FRI *)"

  # If the expression is wrong, an error will occur.
  # value = provider::cronplan::expr("cron(5 0 ? * ? *)")
  #=> ERROR: "'?' cannot be set to both day-of-month and day-of-week."
}
