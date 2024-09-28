package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/winebarrel/terraform-provider-cronplan/internal/expression"
)

type expressionValidator struct {
}

func (v expressionValidator) Description(ctx context.Context) string {
	return "schedule expression must be valid"
}

func (v expressionValidator) MarkdownDescription(ctx context.Context) string {
	return "schedule expression must be valid"
}

func (v expressionValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	if err := expression.Validate(req.ConfigValue.ValueString()); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Schedule Expression",
			fmt.Sprintf("Schedule expression must be valid, got error: %s.", err),
		)

		return
	}
}

func expressionValid() expressionValidator {
	return expressionValidator{}
}
