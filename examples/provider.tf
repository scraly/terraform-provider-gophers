# provider "gophers" {
#   # example configuration here
# }

terraform {
  required_providers {
    gophers = {
      source = "terraform.local/local/gophers"
      version = "0.0.1"
    }
  }
}