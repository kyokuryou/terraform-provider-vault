package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcePublicKey(t *testing.T) {
	t.Skip("data source not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePublicKey,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"data.vault_public_key_data_source.test", "comment", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccDataSourcePublicKey = `
data "vault_public_key_data_source" "test" {
  comment = "bar"
}
`
