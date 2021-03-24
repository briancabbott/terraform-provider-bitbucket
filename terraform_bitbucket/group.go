package terraform_bitbucket



func ResourceGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: GroupCreate,
		ReadContext: GroupRead,
		UpdateContext: GroupUpdate, 
		DeleteContext: GroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		// A Bitbucket project. Projects are used by teams to organize repositories.
		Schema: map[string] *schema.Schema {
			"links": {
				Type: schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: ""
				Elem: &schema.Resource{
					"self": {
						Type: schema.TypeSet,
						Required:    true,
						ForceNew:    true,
						Description: "",
						Elem: &schema.Resource{
							"href": {
								Type: schema.TypeSet,
								Required:    true,
								ForceNew:    true,
								Description: "",
							},
							"name": {
								Type: schema.TypeSet,
								Required:    true,
								ForceNew:    true,
								Description: "",				
							},
						},
					},
					"html": {
						Type: schema.TypeSet,
						Required:    true,
						ForceNew:    true,
						Description: "",
						Elem: &schema.Resource{
							"href": {
								Type: schema.TypeSet,
								Required:    true,
								ForceNew:    true,
								Description: "",
							},
							"name": {
								Type: schema.TypeSet,
								Required:    true,
								ForceNew:    true,
								Description: "",
							},
						},
					},
				},
			},
			"owner": {
				Type: schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: "",
				// "$ref": "#/definitions/account"
			},
			"workspace": {
				Type: schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: "",
				// "$ref": "#/definitions/workspace"
			},
			"name": { 
				Type: schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: "",
			},
			"slug": {
				Type: schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: "The 'sluggified' version of the group's name. This contains only ASCII characters and can therefore be slightly different than the name full_slug"
			}
		}
	}
}


func GroupCreate() *schema.Resource {

}

func GroupRead() *schema.Resource {

}

func GroupUpdate() *schema.Resource {

}

func GroupDelete() *schema.Resource { 

}



// "group": {
// 	"allOf": [
// 		{
// 			"$ref": "#/definitions/object"
// 		},
// 		{
// 			"type": "object",
// 			"title": "Group",
// 			"description": "A group object",
// 			"properties": {
// 				"links": {
// 					"type": "object",
// 					"properties": {
// 						"self": {
// 							"type": "object",
// 							"title": "Link",
// 							"description": "A link to a resource related to this object.",
// 							"properties": {
// 								"href": {
// 									"type": "string",
// 									"format": "uri"
// 								},
// 								"name": {
// 									"type": "string"
// 								}
// 							},
// 							"additionalProperties": false
// 						},
// 						"html": {
// 							"type": "object",
// 							"title": "Link",
// 							"description": "A link to a resource related to this object.",
// 							"properties": {
// 								"href": {
// 									"type": "string",
// 									"format": "uri"
// 								},
// 								"name": {
// 									"type": "string"
// 								}
// 							},
// 							"additionalProperties": false
// 						}
// 					},
// 					"additionalProperties": false
// 				},
// 				"owner": {
// 					"$ref": "#/definitions/account"
// 				},
// 				"workspace": {
// 					"$ref": "#/definitions/workspace"
// 				},
// 				"name": {
// 					"type": "string"
// 				},
// 				"slug": {
// 					"type": "string",
// 					"description": ""
// 				},
// 				"full_slug": {
// 					"type": "string",
// 					"description": "The concatenation of the workspace's slug and the group's slug,\nseparated with a colon (e.g. `acme:developers`)\n"
// 				}
// 			},
// 			"additionalProperties": true
// 		}
// 	]
// },