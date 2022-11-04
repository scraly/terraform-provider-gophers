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
			"displayname": {
				Description: "Display name of a Gopher.",
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
	Name        string `json:"name"`
	DisplayName string `json:"displayname"`
	URL         string `json:"url"`
}

func resourceGopherCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// Retrieve endpoint information
	myClient := meta.(*apiClient)

	//Get the field "name"
	gopherName := d.Get("name").(string)
	gopherDisplayName := d.Get("displayname").(string)
	gopherURL := d.Get("url").(string)

	//Create JSON object with our Gopher
	aGopher := Gopher{
		Name:        gopherName,
		DisplayName: gopherDisplayName,
		URL:         gopherURL,
	}

	//Convert Gopher to byte using json.Marshal method
	body, err := json.Marshal(aGopher)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Will create a Gopher: %+v", string(body))

	// endpoint := fmt.Sprintf("%s/gopher", "http://localhost:8080")
	endpoint := fmt.Sprintf("%s/gopher", myClient.Endpoint)
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
		return diag.Errorf("A Gopher with this name already exits")
	}

	return diags
}

func resourceGopherRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// Retrieve endpoint information
	myClient := meta.(*apiClient)

	//Get the field "name"
	gopherName := d.Get("name").(string)

	log.Printf("[DEBUG] Will read gopher with the name: %s", gopherName)

	// endpoint := fmt.Sprintf("%s/gopher?name=%s", "http://localhost:8080", gopherName)
	endpoint := fmt.Sprintf("%s/gopher?name=%s", myClient.Endpoint, gopherName)
	log.Println("[DEBUG] Endpoint:", endpoint)

	//This function creates a new GET request to localhost:8080/gopher. Then, it decodes the response into a []map[string]interface{}.
	//curl http://localhost:8080/gophers\?name\=yoda-gopher
	resp, err := http.Get(endpoint)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	log.Println("[DEBUG] Status Code", resp.StatusCode)

	if resp.StatusCode == http.StatusOK {
		myGopher := make(map[string]interface{})
		err = json.NewDecoder(resp.Body).Decode(&myGopher)
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

		// set the ID in the TF State
		d.SetId(gopherName)
	} else {
		//When we have 404 HTTP Error Code, returns a warning message in the diagnostic
		return diag.Errorf(" Gopher does not exist")
	}

	return diags
}

//Update the display name and the URL of a gopher
func resourceGopherUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// Retrieve endpoint information
	myClient := meta.(*apiClient)

	//Get the fields
	gopherName := d.Get("name").(string)
	gopherDisplayName := d.Get("displayname").(string)
	gopherURL := d.Get("url").(string)

	//Create JSON object with our Gopher
	aGopher := Gopher{
		Name:        gopherName,
		DisplayName: gopherDisplayName,
		URL:         gopherURL,
	}

	//Convert Gopher to byte using json.Marshal method
	body, err := json.Marshal(aGopher)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Will update a Gopher: %+v", string(body))

	// endpoint := fmt.Sprintf("%s/gopher", "http://localhost:8080")
	endpoint := fmt.Sprintf("%s/gopher", myClient.Endpoint)
	log.Println("[DEBUG] Endpoint:", endpoint)

	//curl -X PUT http://localhost:8080/gopher
	req, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	log.Println("[DEBUG] Status Code", r.StatusCode)

	if r.StatusCode == http.StatusCreated {
		//Unmarshal the body response to the TF State
		gopher := make(map[string]interface{})
		err = json.NewDecoder(r.Body).Decode(&gopher)
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
		return diag.Errorf(" Gopher does not exist")
	}

	return diags
}

func resourceGopherDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// Retrieve endpoint information
	myClient := meta.(*apiClient)

	//Get the field "name"
	gopherName := d.Get("name").(string)

	log.Printf("[DEBUG] Will delete gopher with the name: %s", gopherName)

	// req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/gopher?name=%s", "http://localhost:8080", gopherName), nil)
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/gopher?name=%s", myClient.Endpoint, gopherName), nil)

	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	if r.StatusCode == http.StatusOK {
		//Set ID to empty (Terraform wil remove it from the TF State)
		d.SetId("")
	} else {
		return diag.Errorf(" Gopher does not exist")
	}

	return diags
}

//TODO: import not implemented
// return diag.Errorf("not implemented")
