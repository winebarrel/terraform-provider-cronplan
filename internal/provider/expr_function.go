package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/winebarrel/terraform-provider-cronplan/internal/expression"
)

var (
	_ function.Function = ExprFunction{}
)

func NewExprFunction() function.Function {
	return ExprFunction{}
}

type ExprFunction struct{}

func (r ExprFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "expr"
}

func (r ExprFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "expr function",
		MarkdownDescription: "Valudate the expression and returns it if it is OK.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "expr",
				MarkdownDescription: "Amazon EventBridge schedule expression.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r ExprFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var expr string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &expr))

	if resp.Error != nil {
		return
	}

	_, err := expression.Eval(expr, "", 1)

	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, expr))
}
