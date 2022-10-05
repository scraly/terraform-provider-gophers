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

data "gophers_gopher" "err" {
    name = "dfsdw"
}