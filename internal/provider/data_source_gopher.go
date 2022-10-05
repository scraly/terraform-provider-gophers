package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGopher() *schema.Resource {
	return &schema.Resource{
		Description: "Information about a Gopher.",

		ReadContext: dataSourceGopherRead,

		//{"name":"5th-element","path":"5th-element.png","url":"https://raw.githubusercontent.com/scraly/gophers/main/5th-element.png"}
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			//Computed
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			// "gophers": &schema.Schema{
			// 	Type:     schema.TypeList,
			// 	Computed: true,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"name": &schema.Schema{
			// 				Type:     schema.TypeString,
			// 				Computed: true,
			// 			},
			// 			"path": &schema.Schema{
			// 				Type:     schema.TypeString,
			// 				Computed: true,
			// 			},
			// 			"url": &schema.Schema{
			// 				Type:     schema.TypeString,
			// 				Computed: true,
			// 			},
			// 		},
			// 	},
			// },
		},
	}
}

func dataSourceGopherRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

	// The /gopher?name=%s endpoint returns a gopher object.

	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the field "name"
	gopherName := d.Get("name").(string)

	log.Printf("[DEBUG] Will read gopher with the name: %s", gopherName)

	//This function creates a new GET request to localhost:8080/gopher. Then, it decodes the response into a []map[string]interface{}.
	//curl http://localhost:8080/gophers\?name\=yoda-gopher
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/gopher?name=%s", "http://localhost:8080", gopherName), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	//TODO: when we have 404 HTTP Error Code, returns a warning message in the diagnostic

	myGopher := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&myGopher)
	if err != nil {
		return diag.FromErr(err)
	}

	//Add gopher's information in the TF State the response body
	//assigning each value to its respective schema position.
	for k, v := range myGopher {
		if k != "id" {
			d.Set(k, v)
		} else {
			d.SetId(fmt.Sprint(v))
		}
	}

	// always run
	d.SetId(gopherName)

	//Notice that this function returns a diag.Diagnostics type, which can return multiple errors and warnings to Terraform, giving users more robust error and warning messages.
	return diags
}
