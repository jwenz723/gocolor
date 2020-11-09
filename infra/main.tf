resource "null_resource" "example" {
}

variable "something" {
  default = "default value"
}

output "something-out" {
  value = var.something  
}
