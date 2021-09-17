package provider

import (
	"context"
	"encoding/json"
	"github.com/kyokuryou/terraform-provider-vault/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecret() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "storage data source in the Terraform provider vault.",

		ReadContext: dataSourceSecretRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Default:     "",
				Description: "The name for storage name",
			},
		},
	}
}

func dataSourceSecretRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	var diags diag.Diagnostics
	text, err := meta.(*client.Client).Decode(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(text, &m); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("data", m); err != nil {
		return diag.FromErr(err)
	}
	return diags
}
