# Terraform Provider Scaffolding (Terraform Plugin SDK)



## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.18
-   [gophers API](https://github.com/scraly/gophers-api) running locally

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the `make install` command: 

```bash
$ make install
```

Result:

```bash
go build -o terraform-provider-gophers
mkdir -p ~/.terraform.d/plugins/terraform.local/local/gophers/0.0.1/darwin_amd64
mv terraform-provider-gophers ~/.terraform.d/plugins/terraform.local/local/gophers/0.0.1/darwin_amd64
```

## Using the provider

1. Define the provider you want to use

The provider is defined in `examples/provider.tf` file:

```
terraform {
  required_providers {
    gophers = {
      source  = "terraform.local/local/gophers"
      version = "0.0.1"
    }
  }
}
```

2. Define the datasources

We defined several datasources in `examples/data.f` file:

```yaml
# List of available gophers
data "gophers" "my_gophers" {
}

output "return_gophers" {
  value = length(data.gophers.my_gophers.gophers) >= 1
}

# Display information about a Gopher
data "gophers_gopher" "moultipass" {
  name = "5th-element"
}
```

3. Define resources

We defined a resource in `examples/resource.f` file:

```yaml
resource "gophers_gopher" "x-files" {
  name = "x-files"
  path = "x-files.png"
  url  = "https://raw.githubusercontent.com/scraly/gophers/main/x-files.png"
}
```

4. Init Terraform

```
rm .terraform.lock.hcl && terraform init
```

5. Apply the configuration you defined

```
terraform apply
```

6. destroy datasource and resources you created 

```
terraform destroy
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

# Documentation

When the plugin will be relkeased, add the following content in the `main.tf` file:

```go
// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
```

And execute the `go generate` command to generate the documentation:

```bash
go generate
```