package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func resourceGopherCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the field "name"
	gopherName := d.Get("name").(string)
	gopherPath := d.Get("path").(string)
	gopherURL := d.Get("url").(string)

	//Create JSON object with our Gopher
	aGopher := Gopher{
		Name: gopherName,
		Path: gopherPath,
		URL:  gopherURL,
	}

	//Convert Gopher to byte using json.Marshal method
	body, err := json.Marshal(aGopher)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Will create a Gopher: %+v", string(body))

	endpoint := fmt.Sprintf("%s/gopher", "http://localhost:8080")
	log.Println("[DEBUG] Endpoint:", endpoint)

	//This function creates a new POST request to localhost:8080/gopher.
	//curl http://localhost:8080/gopher
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	log.Println("[DEBUG] Status Code", resp.StatusCode)

	if resp.StatusCode == http.StatusCreated {
		//Unmarshal the body response to the TF State
		gopher := make(map[string]interface{})
		err = json.NewDecoder(resp.Body).Decode(&gopher)
		if err != nil {
			return diag.FromErr(err)
		}

		//Add gopher's information in the TF State the response body
		//assigning each value to its respective schema position.
		for k, v := range gopher {
			if k != "id" {
				d.Set(k, v)
			} else {
				d.SetId(fmt.Sprint(v))
			}
		}

		// Store the Id of the Gopher in the TF State (equals to the gopherName)
		d.SetId(gopherName)

	} else {
		return diag.FromErr(err)
	}

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
