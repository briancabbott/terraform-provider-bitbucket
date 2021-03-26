package terraform_bitbucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	// TODO: log.Printf("[INFO] Creating Provider")
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("", ""),
			},
		},
		ConfigureContextFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"truemark-bitbucket_deployment": resourceDeployment(),
			// "truemark-bitbucket_group": resourceGroup(),
			// "truemark-bitbucket_project": resourceProject(),
			// "truemark-bitbucket_repository": resourceRepository(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// TODO: log.Printf("[INFO] Initializing ConfluentCloud client")
	// username := d.Get("username").(string)
	// password := d.Get("password").(string)

	var diags diag.Diagnostics
	// c := confluentcloud.NewClient(username, password)
	// loginErr := c.Login()
	// if loginErr == nil {
		return nil, diags
	// }
	// err := resource.RetryContext(ctx, 30*time.Minute, func() *resource.RetryError {
	// 	err := c.Login()
	// 	if strings.Contains(err.Error(), "Exceeded rate limit") {
	// 		// TODO: log.Printf("[INFO] ConfluentCloud API rate limit exceeded, retrying.")
	// 		return resource.RetryableError(err)
	// 	}
	// 	return resource.NonRetryableError(err)
	// })
	// return c, diag.FromErr(err)
}

func resourceDeployment() *schema.Resource {
	return &schema.Resource{

	}
}

func resourceGroup() *schema.Resource {
	return &schema.Resource{

	}
}

func resourceProject() *schema.Resource {
	return &schema.Resource{

	}
}

func resourceRepository() *schema.Resource {
	return &schema.Resource{

	}
}
