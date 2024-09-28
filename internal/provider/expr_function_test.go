package provider_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestExprFunction_valid(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::cronplan::expr("cron(5 0 ? * MON-FRI *)")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("test", "cron(5 0 ? * MON-FRI *)"),
				),
			},
		},
	})
}

func TestExprFunction_invalid(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::cronplan::expr("cron(5 0 ? * ? *)")
				}
				`,
				ExpectError: regexp.MustCompile(strings.ReplaceAll(` `, `\s+`, regexp.QuoteMeta(`Failed to parse expr: 'cron(5 0 ? * ? *)': '?' cannot be set to both day-of-month and day-of-week.`))),
			},
		},
	})
}
