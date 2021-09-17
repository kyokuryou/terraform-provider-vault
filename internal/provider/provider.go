package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kyokuryou/terraform-provider-vault/client"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"path": {
					Type:        schema.TypeString,
					Optional:    false,
					ForceNew:    true,
					Description: "The path to the save storage account directory",
					DefaultFunc: schema.EnvDefaultFunc("VAULT_PATH", ""),
				},
				"private_key": {
					Type:        schema.TypeString,
					Optional:    false,
					ForceNew:    true,
					Description: "The Private Key which should be used for authentication, which needs to rsa format",
					DefaultFunc: schema.EnvDefaultFunc("VAULT_PRIVATE_KEY", ""),
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				"vault_secret_resource": resourceSecret(),
			},
			DataSourcesMap: map[string]*schema.Resource{
				"vault_secret_data_source":     dataSourceSecret(),
				"vault_public_key_data_source": dataSourcePublicKey(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		p.UserAgent("terraform-provider-vault", version)
		c, err := client.New(d.Get("path").(string), d.Get("private_key").(string))
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return &c, nil
	}
}
