package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shihanng/terraform-provider-installer/internal/provider"
	"gotest.tools/v3/assert"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){ //nolint:gochecknoglobals
	"installer": func() (*schema.Provider, error) { //nolint:unparam
		return provider.New("dev")(), nil
	},
}

func TestProvider(t *testing.T) {
	t.Parallel()

	t.Run("runs internal validation", func(t *testing.T) {
		t.Parallel()

		assert.NilError(t, provider.New("dev")().InternalValidate())
	})

	// No sure if this has any value but we can find this pattern in:
	// - https://github.com/integrations/terraform-provider-github/blob/f9508c5a4012e25400853bbb684877e3f991268f/github/provider_test.go#L48
	// - https://github.com/hashicorp/terraform-provider-hashicups/blob/a7e659e5551b717b268ca64c901e255ed6ed55e5/hashicups/provider_test.go#L27
	// nolint:lll
	t.Run("has an implementation", func(t *testing.T) {
		t.Parallel()

		_ = provider.New("dev")()
	})
}

func testAccPreCheck(t *testing.T) { //nolint:thelper
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}
