package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceStorage(t *testing.T) {
	t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceStorage,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"vault_secret_resource.test", "name", regexp.MustCompile("^te")),
				),
			},
		},
	})
}

const testAccResourceStorage = `
resource "vault_secret_resource" "test" {
  name = "test"
  vars = {
     ip = "11111"
     username = "2222"
     password = "33333"
  }
}
`
