package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGopher() *schema.Resource {
	return &schema.Resource{
		Description: "Information about a Gopher.",

		ReadContext: dataSourceGopherRead,

		//[{"name":"5th-element","path":"5th-element.png","url":"https://raw.githubusercontent.com/scraly/gophers/main/5th-element.png"}]
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"gophers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"path": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGopherRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

	client := &http.Client{Timeout: 10 * time.Second}

	// The /gophers endpoint returns an array of gophers.

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the field "name"
	gopherName := d.Get("name").(string)

	//This function creates a new GET request to localhost:8080/gophers. Then, it decodes the response into a []map[string]interface{}.
	//curl http://localhost:8080/gophers\?name\=yoda-gopher
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/gophers?name=%s", "http://localhost:8080", gopherName), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	gophers := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&gophers)
	if err != nil {
		return diag.FromErr(err)
	}

	//The d.Set("gophers", gophers) function sets the response body (list of coffees object) to Terraform coffees data source, assigning each value to its respective schema position.
	if err := d.Set("gophers", gophers); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	//Notice that this function returns a diag.Diagnostics type, which can return multiple errors and warnings to Terraform, giving users more robust error and warning messages.
	return diags
}
