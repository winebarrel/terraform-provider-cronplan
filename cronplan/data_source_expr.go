package cronplan

import (
	"context"
	"time"

	"github.com/araddon/dateparse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cp "github.com/winebarrel/cronplan"
)

func dataSourceExpr() *schema.Resource {
	return &schema.Resource{
		ReadContext: readExpr,
		Schema: map[string]*schema.Schema{
			"cron": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val any, key string) (warns []string, errs []error) {
					if _, err := cp.Parse(val.(string)); err != nil {
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
	expr := d.Get("cron").(string)
	from := d.Get("from").(string)
	n := d.Get("num_schedules").(int)
	schedule, err := evalCron(expr, from, n)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id.UniqueId())
	d.Set("schedules", schedule) //nolint:errcheck

	return nil
}

func evalCron(expr string, fromStr string, n int) ([]string, error) {
	cron, err := cp.Parse(expr)

	if err != nil {
		return nil, err
	}

	from := time.Now()

	if fromStr != "" {
		from, err = dateparse.ParseAny(fromStr)

		if err != nil {
			return nil, err
		}
	}

	schedule := []string{}
	next := cron.NextN(from, n)

	for _, v := range next {
		schedule = append(schedule, v.Format("Mon, 02 Jan 2006 15:04:05"))
	}

	return schedule, nil
}
