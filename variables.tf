variable "name" {
  description = "The name of the resource"
  type        = string
}

variable "display_name" {
  description = "The display name for the management group"
  type        = string
}

variable "parent_management_group_id" {
  description = "The parent ID for the management group"
  type        = string
  default     = ""
}

variable "subscription_ids" {
  description = "set of subscription IDs to associate with the management group"
  type        = set(string)
}
