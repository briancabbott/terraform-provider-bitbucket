package terraform_bitbucket



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
			"truemark-bitbucket_group": resourceGroup(),
			"truemark-bitbucket_project": resourceProject(),
			"truemark-bitbucket_repository": resourceRepository(),
		},
	}
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
