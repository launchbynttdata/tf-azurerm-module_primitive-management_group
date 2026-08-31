// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

variable "resource_names_map" {
  description = "Map of resource names used by tf-launch-module_library-resource_name"
  type = map(object({
    name       = string
    max_length = optional(number, 80)
  }))
}

variable "logical_product_family" {
  type        = string
  description = "Product family for generated resource names"
  default     = "launch"
}

variable "logical_product_service" {
  type        = string
  description = "Product service for generated resource names"
  default     = "gotest"
}

variable "class_env" {
  type        = string
  description = "Environment class for generated resource names"
  default     = "sandbox"
}

variable "instance_env" {
  type        = number
  description = "Instance number for the environment"
  default     = 0
}

variable "instance_resource" {
  type        = number
  description = "Instance number for the resource"
  default     = 0
}

variable "region" {
  type        = string
  description = "Azure region token used by the resource naming module"
  default     = "eastus"
}

variable "spoke_subscription_ids" {
  type        = set(string)
  description = "Subscription IDs for the spoke subscriptions"
}
