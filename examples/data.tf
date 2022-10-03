# data "scaffolding_data_source" "example" {
#   sample_attribute = "foo"
# }

# List of available gophers
data "gophers" "my_gophers" {
}

output "return_gophers" {
	value = length(data.gophers.my_gophers.gophers) >= 1
}