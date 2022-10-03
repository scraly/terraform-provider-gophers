package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGopher(t *testing.T) {
	resourceName := "data.gophers_gopher.yoda"

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGopher,
				Check: resource.ComposeTestCheckFunc(

					resource.TestCheckResourceAttrSet(resourceName, "gophers.#"),
				),
			},
		},
	})
}

const testAccDataSourceGopher = `
data "gophers_gopher" "yoda" {
    name = "yoda-gopher"
}
`
