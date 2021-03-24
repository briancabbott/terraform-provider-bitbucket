package terraform_bitbucket




///
// Deployment Schema:
//    - uuid
//    - state
//    - environment
//    - release

func ResourceDeployment() *schema.Resource {
	return &schema.Resource{
		CreateContext: DeploymentCreate,
		ReadContext: DeploymentRead,
		UpdateContext: DeploymentUpdate, 
		DeleteContext: DeploymentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		// A Bitbucket project. Projects are used by teams to organize repositories.
		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString,
				Description: "The UUID identifying the deployment."
			},
			"state": {
				"$ref": "#/definitions/deployment_state"
				"deployment_state": {
					"allOf": [
						{
							"$ref": "#/definitions/object"
						},
						{
							"additionalProperties": true,
							"type": "object",
							"description": "The representation of the progress state of a deployment.",
							"properties": {}
						}
					]
				},

				// "deployment_state_completed_status": {
				// 	"allOf": [
				// 		{
				// 			"$ref": "#/definitions/object"
				// 		},
				// 		{
				// 			"additionalProperties": true,
				// 			"type": "object",
				// 			"description": "The status of a completed deployment.",
				// 			"properties": {}
				// 		}
				// 	]
				// },
				// "deployment_state_in_progress": {
				// 	"allOf": [
				// 		{
				// 			"$ref": "#/definitions/deployment_state"
				// 		},
				// 		{
				// 			"additionalProperties": true,
				// 			"type": "object",
				// 			"description": "A Bitbucket Deployment IN_PROGRESS deployment state.",
				// 			"properties": {
				// 				"name": {
				// 					"enum": [
				// 						"IN_PROGRESS"
				// 					],
				// 					"type": "string",
				// 					"description": "The name of deployment state (IN_PROGRESS)."
				// 				},
				// 				"url": {
				// 					"type": "string",
				// 					"format": "uri",
				// 					"description": "Link to the deployment result."
				// 				},
				// 				"deployer": {
				// 					"$ref": "#/definitions/account",
				// 					"description": "The Bitbucket account that was used to perform the deployment."
				// 				},
				// 				"start_date": {
				// 					"type": "string",
				// 					"format": "date-time",
				// 					"description": "The timestamp when the deployment was started."
				// 				}
				// 			}
				// 		}
				// 	]
				// },
				// "deployment_state_undeployed": {
				// 	"allOf": [
				// 		{
				// 			"$ref": "#/definitions/deployment_state"
				// 		},
				// 		{
				// 			"additionalProperties": true,
				// 			"type": "object",
				// 			"description": "A Bitbucket Deployment UNDEPLOYED deployment state.",
				// 			"properties": {
				// 				"name": {
				// 					"enum": [
				// 						"UNDEPLOYED"
				// 					],
				// 					"type": "string",
				// 					"description": "The name of deployment state (UNDEPLOYED)."
				// 				},
				// 				"trigger_url": {
				// 					"type": "string",
				// 					"format": "uri",
				// 					"description": "Link to trigger the deployment."
				// 				}
				// 			}
				// 		}
				// 	]
				// },
				// "deployment_state_completed_status_stopped": {
				// 	"allOf": [
				// 		{
				// 			"$ref": "#/definitions/deployment_state_completed_status"
				// 		},
				// 		{
				// 			"additionalProperties": true,
				// 			"type": "object",
				// 			"description": "A STOPPED completed deployment status.",
				// 			"properties": {
				// 				"name": {
				// 					"enum": [
				// 						"STOPPED"
				// 					],
				// 					"type": "string",
				// 					"description": "The name of the completed deployment status (STOPPED)."
				// 				}
				// 			}
				// 		}
				// 	]
				// },
				// "deployment_state_completed_status_failed": {
				// 	"allOf": [
				// 		{
				// 			"$ref": "#/definitions/deployment_state_completed_status"
				// 		},
				// 		{
				// 			"additionalProperties": true,
				// 			"type": "object",
				// 			"description": "A FAILED completed deployment status.",
				// 			"properties": {
				// 				"name": {
				// 					"enum": [
				// 						"FAILED"
				// 					],
				// 					"type": "string",
				// 					"description": "The name of the completed deployment status (FAILED)."
				// 				}
				// 			}
				// 		}
				// 	]
				// },
			},
			"environment": {
				"$ref": "#/definitions/deployment_environment",
				"description": "A deployment environment."
				// "deployment_environment": {
				// 	"allOf": [
				// 		{
				// 			"$ref": "#/definitions/object"
				// 		},
				// 		{
				// 			"additionalProperties": true,
				// 			"type": "object",
				// 			"description": "A Bitbucket Deployment Environment.",
				// 			"properties": {
				// 				"uuid": {
				// 					"type": "string",
				// 					"description": "The UUID identifying the environment."
				// 				},
				// 				"name": {
				// 					"type": "string",
				// 					"description": "The name of the environment."
				// 				}
				// 			}
				// 		}
				// 	],
				// 	"x-bb-default-fields": [
				// 		"uuid"
				// 	],
				// 	"x-bb-url": "/rest/2.0/accounts/{target_user.uuid}/repositories/{repository.uuid}/environments/{uuid}"
				// },
			},
			"release": {
				"$ref": "#/definitions/deployment_release",
				"description": "The release associated with this deployment."
			}
			"allOf": [
                {
                    "$ref": "#/definitions/object"
                },
                {
                    "additionalProperties": true,
                    "type": "object",
                    "description": "A Bitbucket Deployment.",
                    "properties": {
                        
                    }
                }
            ]
        },
	}
}


