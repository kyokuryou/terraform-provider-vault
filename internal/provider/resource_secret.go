package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kyokuryou/terraform-provider-vault/client"
	"strings"
)

func resourceSecret() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description:   "secret resource in the Terraform provider vault.",
		CreateContext: resourceSecretCreate,
		ReadContext:   resourceSecretRead,
		UpdateContext: resourceSecretUpdate,
		DeleteContext: resourceSecretDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Path to the directory where the files to template reside",
				Required:    true,
				ForceNew:    true,
			},
			"vars": {
				Type:         schema.TypeMap,
				Optional:     true,
				Default:      make(map[string]interface{}),
				Description:  "Variables to substitute",
				ValidateFunc: validateVarsAttribute,
				ForceNew:     true,
			},
		},
	}
}

func resourceSecretCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	b, err := json.Marshal(d.Get("vars").(map[string]interface{}))
	if err != nil {
		return diag.FromErr(err)
	}
	if id, e := meta.(*client.Client).Sign(b); e == nil {
		d.SetId(id)
	}
	if e := meta.(*client.Client).Encode(d.Get("name").(string), b); e != nil {
		return diag.FromErr(e)
	}
	return diags
}
func resourceSecretUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	b, err := json.Marshal(d.Get("vars").(map[string]interface{}))
	if err != nil {
		return diag.FromErr(err)
	}
	if e := meta.(*client.Client).Encode(d.Get("name").(string), b); e != nil {
		return diag.FromErr(e)
	}
	return diags
}

func resourceSecretRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceSecretDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	name := d.Get("name").(string)
	if e := meta.(*client.Client).Delete(name); e != nil {
		return diag.FromErr(e)
	}
	return diags
}

func validateVarsAttribute(v interface{}, key string) (ws []string, es []error) {
	// vars can only be primitives right now
	var badVars []string
	for k, v := range v.(map[string]interface{}) {
		switch v.(type) {
		case []interface{}:
			badVars = append(badVars, fmt.Sprintf("%s (list)", k))
		case map[string]interface{}:
			badVars = append(badVars, fmt.Sprintf("%s (map)", k))
		}
	}
	if len(badVars) > 0 {
		es = append(es, fmt.Errorf(
			"%s: cannot contain non-primitives; bad keys: %s",
			key, strings.Join(badVars, ", ")))
	}
	return
}
