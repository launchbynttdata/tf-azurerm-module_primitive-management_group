<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | ~> 3.77 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_management_group"></a> [management\_group](#module\_management\_group) | ../.. | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_management_group"></a> [management\_group](#input\_management\_group) | Management group values | <pre>object({<br/>    name         = string<br/>    display_name = string<br/>  })</pre> | n/a | yes |
| <a name="input_spoke_subscription_ids"></a> [spoke\_subscription\_ids](#input\_spoke\_subscription\_ids) | Subscription IDs for the spoke subscriptions | `set(string)` | n/a | yes |

## Outputs

No outputs.
<!-- END_TF_DOCS -->
