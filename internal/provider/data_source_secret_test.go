package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceStorage(t *testing.T) {
	t.Skip("data source not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceStorage,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.vault_secret_data_source.test", "data", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccDataSourceStorage = `
data "vault_secret_data_source" "test" {
  name = "bar"
}
`
