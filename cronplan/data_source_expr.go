package cronplan

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/winebarrel/terraform-provider-cronplan/internal/expression"
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

func dataSourceExpr() *schema.Resource {
	return &schema.Resource{
		ReadContext: readExpr,
		Schema: map[string]*schema.Schema{
			"expr": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val any, key string) (warns []string, errs []error) {
					if err := expression.Validate(val.(string)); err != nil {
						errs = append(errs, err)
					}
					return
				},
			},
			"from": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_schedules": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"schedules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readExpr(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	expr := d.Get("expr").(string)
	from := d.Get("from").(string)
	n := d.Get("num_schedules").(int)
	schedule, err := expression.Eval(expr, from, n)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id.UniqueId())
	d.Set("schedules", schedule) //nolint:errcheck

	return nil
}
