terraform {
  required_providers {
    gophers = {
      source  = "terraform.local/local/gophers"
      version = "0.0.1"
    }
  }
}

provider "gophers" {
  # Configuration options
  # endpoint = "http://localhost:8080"
  endpoint = "https://8080-scraly-gophersapi-pdiocn3y9uh.ws-eu73.gitpod.io"
}