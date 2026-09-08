<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | ~> 3.77 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_management_group"></a> [management\_group](#module\_management\_group) | ../.. | n/a |
| <a name="module_resource_names"></a> [resource\_names](#module\_resource\_names) | terraform.registry.launch.nttdata.com/module_library/resource_name/launch | ~> 2.4 |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_class_env"></a> [class\_env](#input\_class\_env) | Environment class for generated resource names | `string` | `"sandbox"` | no |
| <a name="input_instance_env"></a> [instance\_env](#input\_instance\_env) | Instance number for the environment | `number` | `0` | no |
| <a name="input_instance_resource"></a> [instance\_resource](#input\_instance\_resource) | Instance number for the resource | `number` | `0` | no |
| <a name="input_logical_product_family"></a> [logical\_product\_family](#input\_logical\_product\_family) | Product family for generated resource names | `string` | `"launch"` | no |
| <a name="input_logical_product_service"></a> [logical\_product\_service](#input\_logical\_product\_service) | Product service for generated resource names | `string` | `"gotest"` | no |
| <a name="input_region"></a> [region](#input\_region) | Azure region token used by the resource naming module | `string` | `"eastus"` | no |
| <a name="input_resource_names_map"></a> [resource\_names\_map](#input\_resource\_names\_map) | Map of resource names used by tf-launch-module\_library-resource\_name | <pre>map(object({<br/>    name       = string<br/>    max_length = optional(number, 80)<br/>  }))</pre> | n/a | yes |
| <a name="input_spoke_subscription_ids"></a> [spoke\_subscription\_ids](#input\_spoke\_subscription\_ids) | Subscription IDs for the spoke subscriptions | `set(string)` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_management_group_id"></a> [management\_group\_id](#output\_management\_group\_id) | n/a |
| <a name="output_management_group_name"></a> [management\_group\_name](#output\_management\_group\_name) | n/a |
<!-- END_TF_DOCS -->
