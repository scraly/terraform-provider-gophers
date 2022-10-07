package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceGopher(t *testing.T) {

	resourceName := "gophers_gopher.gandalf"

	name := "gandalf"
	path := "gandalf.png"
	url := "https://raw.githubusercontent.com/scraly/gophers/main/gandalf-colored.png"
	config := fmt.Sprintf(
		testAccResourceGopher,
		name,
		path,
		url,
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "path", path),
					resource.TestCheckResourceAttr(resourceName, "url", url),
				),
			},
		},
	})
}

const testAccResourceGopher = `
resource "gophers_gopher" "gandalf" {
	name = "%s"
	path = "%s"
	url  = "%s"
}
`
