package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGopher() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Gopher resource in the Terraform provider gopher.",

		CreateContext: resourceGopherCreate,
		ReadContext:   resourceGopherRead,
		UpdateContext: resourceGopherUpdate,
		DeleteContext: resourceGopherDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of a Gopher.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"path": {
				Description: "Path of a Gopher.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"url": {
				Description: "URL of a Gopher.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

type Gopher struct {
	Name string `json:"name"`
	Path string `json:"path"`
	URL  string `json:"url"`
}

//TODO:
func resourceGopherCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the field "name"
	gopherName := d.Get("name").(string)
	gopherPath := d.Get("path").(string)
	gopherURL := d.Get("url").(string)

	//Create JSON object with our Gopher
	myGopher, err := json.Marshal(Gopher{Name: gopherName, Path: gopherPath, URL: gopherURL})
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Will create a Gopher: %+v", string(myGopher))

	bodyReader := bytes.NewReader(myGopher)

	//This function creates a new GET request to localhost:8080/gopher. Then, it decodes the response into a []map[string]interface{}.
	//curl http://localhost:8080/gopher
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/gopher", "http://localhost:8080"), bodyReader)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	//TODO:
	//gopher := make(map[string]interface{})
	gopher := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&gopher)
	if err != nil {
		return diag.FromErr(err)
	}

	//The d.Set("gophers", gophers) function sets the response body (created gopher object) to Terraform gophers_gopher data source, assigning each value to its respective schema position.
	if err := d.Set("gopher", gopher); err != nil {
		return diag.FromErr(err)
	}

	//Store the Id of the Gopher in the TF State (equals to the gopherName)
	d.SetId(gopherName)

	return diags
}

//TODO:
func resourceGopherRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics

	// gopherName := d.Id()
	//TODO: get gopher by name

	//TODO: comme le datasource
	//TODO: recuperer la reponse et setter le d.Set... (name, path, url)

	// return diags

	return diag.Errorf("not implemented")
}

//TODO:
func resourceGopherUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	//TODO: PUT?

	return diag.Errorf("not implemented")
}

//TODO:
func resourceGopherDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	//TODO: Call DELETE API
	//TODO: Set ID a empty (il enlevera du TF state)

	return diag.Errorf("not implemented")
}

//TODO: import not implemented
