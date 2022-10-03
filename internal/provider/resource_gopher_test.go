package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceGopher(t *testing.T) {
	//TODO:
	t.Skip("resource not yet implemented, remove this once you add your own code")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceScaffolding,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"gopher.my_gopher", "name", regexp.MustCompile("^to")),
				),
			},
		},
	})
}

const testAccResourceScaffolding = `
resource "gopher" "my_gopher" {
  name = "toto"
  path = "tutu"
  url = "titi"
}
`
