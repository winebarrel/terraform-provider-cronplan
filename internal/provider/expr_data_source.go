package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/winebarrel/terraform-provider-cronplan/internal/expression"
)

const (
	DefaultNumSchedules int = 10
)

var _ datasource.DataSource = &ExprDataSource{}

func NewExprDataSource() datasource.DataSource {
	return &ExprDataSource{}
}

type ExprDataSource struct {
}

type ExprDataSourceModel struct {
	Expr         types.String   `tfsdk:"expr"`
	From         types.String   `tfsdk:"from"`
	NumSchedules types.Int32    `tfsdk:"num_schedules"`
	Schedules    []types.String `tfsdk:"schedules"`
}

func (d *ExprDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_expr"
}

func (d *ExprDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"expr": schema.StringAttribute{
				MarkdownDescription: "Amazon EventBridge schedule expression. see https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-scheduled-rule-pattern.html",
				Required:            true,
				Validators: []validator.String{
					expressionValid(),
				},
			},
			"from": schema.StringAttribute{
				MarkdownDescription: "Cron expression start date.",
				Optional:            true,
			},
			"num_schedules": schema.Int32Attribute{
				MarkdownDescription: "Number of schedules to output.",
				Optional:            true,
				Validators: []validator.Int32{
					int32validator.AtLeast(0),
				},
			},
			"schedules": schema.ListAttribute{
				MarkdownDescription: "Cron expression schedules",
				Computed:            true,
				ElementType:         types.StringType,
			},
		},
	}
}

func (d *ExprDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *ExprDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ExprDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	expr := data.Expr.ValueString()
	from := ""

	if !data.From.IsUnknown() && !data.From.IsNull() {
		from = data.From.ValueString()
	}

	n := DefaultNumSchedules

	if !data.NumSchedules.IsUnknown() && !data.NumSchedules.IsNull() {
		n = int(data.NumSchedules.ValueInt32())
	}

	schedules, err := expression.Eval(expr, from, n)

	if err != nil {
		resp.Diagnostics.AddError("Eval Error", err.Error())
		return
	}

	for _, skd := range schedules {
		data.Schedules = append(data.Schedules, types.StringValue(skd))
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
