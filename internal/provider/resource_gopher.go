package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//TODO:
func resourceGopher() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider gopher.",

		CreateContext: resourceGopherCreate,
		ReadContext:   resourceGopherRead,
		UpdateContext: resourceGopherUpdate,
		DeleteContext: resourceGopherDelete,

		//TODO:
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

//TODO:
func resourceGopherCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Trace(ctx, "created a resource")

	return diag.Errorf("not implemented")
}

//TODO:
func resourceGopherRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

//TODO:
func resourceGopherUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

//TODO:
func resourceGopherDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
