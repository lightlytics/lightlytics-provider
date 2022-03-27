package lightlytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LIGHTLYTICS_HOST", nil),
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("LIGHTLYTICS_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("LIGHTLYTICS_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"lightlytics_account": resourceAccount(),
		},
		DataSourcesMap: map[string]*schema.Resource{
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := d.Get("host").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	var diags diag.Diagnostics

	if (host != "") && (username != "") && (password != "") {
		c, err := NewClient(&host, &username, &password)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create Lightlytics client",
				Detail:   "Unable to auth user for authenticated Lightlytics client",
			})
			return nil, diags
		}

		return c, diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to create Lightlytics client",
		Detail:   "Missing host, username or password",
	  })

	return nil, diags
}

func marshalData(d *schema.ResourceData, vals map[string]interface{}) {
	for k, v := range vals {
		if k == "id" {
			d.SetId(v.(string))
		} else {
			str, ok := v.(string)
			if ok {
				d.Set(k, str)
			} else {
				d.Set(k, v)
			}
		}
	}
}
