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

`examples/provider.tf`:

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

1. Init Terraform

```
$ rm .terraform.lock.hcl && terraform init
```

1. Define datasources

`examples/data.tf`

```
$ terraform apply
```

1. Define resource

`examples/resource.tf`

```
$ terraform apply
```

1. destroy datasource and resources you created 

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
