package cronplan_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/terraform-provider-cronplan/cronplan"
)

var (
	testProviders map[string]*schema.Provider
	testProvider  *schema.Provider
)

func init() {
	testProvider = cronplan.Provider()
	testProviders = map[string]*schema.Provider{
		"cronplan": testProvider,
	}
}

func TestProvider(t *testing.T) {
	assert := assert.New(t)
	provider := cronplan.Provider()
	err := provider.InternalValidate()
	assert.NoError(err)
}
