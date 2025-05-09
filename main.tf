resource "azurerm_management_group" "management_group" {
  name                       = var.name
  display_name               = var.display_name
  subscription_ids           = var.subscription_ids
  parent_management_group_id = var.parent_management_group_id != "" ? var.parent_management_group_id : null
}
