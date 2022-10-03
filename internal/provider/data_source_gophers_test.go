package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGophers(t *testing.T) {
	resourceName := "data.gophers.my_gophers"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGophers,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(resourceName, "gophers.#"),
					resource.TestCheckOutput("return_gophers", "true"),
				),
			},
		},
	})
}

const testAccDataSourceGophers = `
data "gophers" "my_gophers" {
}

output "return_gophers" {
	value = length(data.gophers.my_gophers.gophers) >= 1
}
`
