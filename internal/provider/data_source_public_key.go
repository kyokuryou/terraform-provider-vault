package provider

import (
	"context"
	"github.com/kyokuryou/terraform-provider-vault/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePublicKey() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "public key data source in the Terraform provider vault.",

		ReadContext: dataSourcePublicKeyRead,

		Schema: map[string]*schema.Schema{
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The comment for a key data",
			},
		},
	}
}

func dataSourcePublicKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	var diags diag.Diagnostics
	pubkey, err := meta.(*client.Client).GenKey(d.Get("comment").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("public_key", pubkey); err != nil {
		return diag.FromErr(err)
	}
	return diags
}
